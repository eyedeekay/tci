package main

import (
	"fmt"
	"github.com/eyedeekay/tci/travis"
	"strings"
)

func init() {
	cmds["log"] = cmd{log, "", "display information about the latest build"}
	cmdHelp["log"] = `Shows a summary of the latest build.
`
}

func log() {
	slug := detectSlug()
	client := travis.NewClient()

	repoResp, _ := client.GetRepository(slug)
	if repoResp == (travis.RepositoryResponse{}) {
		println("Couldn't find repository")
		return
	}

	buildResp, _ := client.GetBuild(repoResp.Repository.LastBuildID)
	if buildResp == (travis.BuildResponse{}) {
		println("Couldn't find build.")
		return
	}

	build := buildResp.Build
	commit := buildResp.Commit

	fmt.Printf(bold("Build #%s: %s\n"), build.Number, strings.Split(commit.Message, "\n")[0])
	printInfo("State", build.State)
    printInfo("Job ID", build.Jobs.Id)
	if build.PullRequest {
		printInfo("Type", "pull request")
	} else {
		printInfo("Type", "push")
	}
	printInfo("Branch", commit.Branch)
	printInfo("Compare URL", commit.CompareURL)
	printInfo("Duration", formatDuration(build.Duration))
	printInfo("Started", formatTime(build.StartedAt))
	printInfo("Finished", formatTime(build.FinishedAt))
}

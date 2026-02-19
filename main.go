package main

import (
	"PReQual/client"
	"PReQual/helper"
	"PReQual/metric"
	"fmt"
	"strings"
)

const workspace = "tmp"
const repo = "ReViSE-EuroSpaceCenter/ReViSE-backend"

func main() {
	var prClient client.PullRequestClient
	prClient = &client.GhClient{}

	var analyzer metric.ProjectAnalyser
	analyzer = &metric.SonarQubeAnalyzer{}

	prs, err := prClient.GetPullRequests(repo)
	if err != nil {
		fmt.Printf("Error fetching pull requests: %v\n", err)
		return
	}

	for _, pr := range prs {
		fmt.Printf("PR #%d: %s (Base: %s, Head: %s)\n", pr.Number, pr.Title, pr.BaseRefOid, pr.HeadRefOid)

		var path = fmt.Sprintf("%s/%s/pr_%d", workspace, repo, pr.Number)

		if err := prClient.RetrieveBranchZip(repo, pr.HeadRefOid, path, "head.zip"); err != nil {
			return
		}
		if err = prClient.RetrieveBranchZip(repo, pr.BaseRefOid, path, "base.zip"); err != nil {
			return
		}

		helper.WriteMetaDataFile(path, pr)

		formattedRepo := strings.Replace(repo, "/", "-", -1)

		err := analyzer.AnalyzeProject(formattedRepo, path)
		if err != nil {
			fmt.Printf("Error analyzing pull requests: %v\n", err)
			return
		}
	}
}

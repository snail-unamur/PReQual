package main

import (
	"PReQual/client"
	"fmt"
)

const workspace = "tmp"

func main() {
	repo := "ReViSE-EuroSpaceCenter/ReViSE-backend"

	var prClient client.PullRequestClient
	prClient = &client.GhClient{}

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
	}
}

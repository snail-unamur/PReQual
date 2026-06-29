package client

import "PReQual/model"

type PullRequestClient interface {
	GetPullRequests(repo string) ([]model.PullRequest, error)
	RetrieveBranchZip(repo, sha, outputPath, outputName string) error
}

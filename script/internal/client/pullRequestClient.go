package client

import (
	"PReQual/internal/domain"
)

type PullRequestClient interface {
	GetPullRequests(repo string) ([]domain.PullRequest, error)
	RetrieveBranchZip(repo, sha, outputPath, outputName string) error
}

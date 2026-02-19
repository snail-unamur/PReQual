package client

import (
	"PReQual/helper"
	"PReQual/model"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

type GhClient struct{}

const limit = 1000
const data = "number,title,baseRefOid,headRefOid,state,createdAt,closedAt,comments,body,closingIssuesReferences"

func (c *GhClient) GetPullRequests(repo string) ([]model.PullRequest, error) {
	args := []string{
		"pr", "list",
		"-R", repo,
		"--state", "all",
		"--limit", strconv.Itoa(limit),
		"--json", data,
	}

	output, err := c.runGh(args)
	if err != nil {
		return nil, fmt.Errorf("get pull requests for %s: %w", repo, err)
	}

	var prs []model.PullRequest
	if err := json.Unmarshal(output, &prs); err != nil {
		return nil, fmt.Errorf("decode pull requests JSON: %w", err)
	}

	return prs, nil
}

func (c *GhClient) RetrieveBranchZip(repo, sha, outputPath, outputName string) error {
	args := []string{
		"api",
		fmt.Sprintf("repos/%s/zipball/%s", repo, sha),
		"--header", "Accept: application/vnd.github+json",
	}

	output, err := c.runGh(args)
	if err != nil {
		return fmt.Errorf("download zip for %s@%s: %w", repo, sha, err)
	}

	if err := helper.SaveToFile(outputPath, outputName, output); err != nil {
		return fmt.Errorf("save zip to %s: %w", outputPath, err)
	}

	return nil
}

func (c *GhClient) runGh(args []string) ([]byte, error) {
	cmd := exec.Command("gh", args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("gh %v failed: %s", args, stderr.String())
	}

	return output, nil
}

package model

type PullRequest struct {
	Number     int    `json:"number"`
	Title      string `json:"title"`
	BaseRefOid string `json:"baseRefOid"`
	HeadRefOid string `json:"headRefOid"`
}

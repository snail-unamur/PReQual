package domain

type Author struct {
	Login string `json:"login"`
	IsBot bool   `json:"is_bot"`
}

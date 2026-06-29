package model

type RateLimitResponse struct {
	Rate struct {
		Remaining int   `json:"remaining"`
		Reset     int64 `json:"reset"`
	} `json:"rate"`
}

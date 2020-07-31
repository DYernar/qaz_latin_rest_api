package model

type SigninResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
	User   User   `json:"user"`
}

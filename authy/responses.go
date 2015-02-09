package authy

type NewUserReq struct {
	User *AuthyUser `json:"user"`
}

type NewUserResp struct {
	User    *AuthyUser        `json:"user"`
	Errors  map[string]string `json:"errors",omitempty`
	Message string            `json:"message",omitempty`
	Success bool              `json:"success",omitempty`
}

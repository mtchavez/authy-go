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

type VerifyReq struct {
	Token        string `json:"token"`
	AuthyId      int    `json:"authy_id"`
	Force        bool   `json:"-"`
	CustomAction string `json:"-"`
}

type VerifyResp struct {
	Token   string            `json:"token"`
	Errors  map[string]string `json:"errors",omitempty`
	Message string            `json:"message",omitempty`
	Success bool              `json:"success",omitempty`
}

type SMSReq struct {
	AuthyId      int    `json:"authy_id"`
	Force        bool   `json:"-"`
	Shortcode    bool   `json:"-"`
	CustomAction string `json:"-"`
}

type SMSResp struct {
	Errors    map[string]string `json:"errors",omitempty`
	Message   string            `json:"message",omitempty`
	Cellphone string            `json:"cellphone",omitempty`
	Success   bool              `json:"success",omitempty`
}

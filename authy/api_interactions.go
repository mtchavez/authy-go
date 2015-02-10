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

type CallReq struct {
	AuthyId int  `json:"authy_id"`
	Force   bool `json:"-"`
}

type CallResp struct {
	Errors  map[string]string `json:"errors",omitempty`
	Message string            `json:"message",omitempty`
	Success bool              `json:"success",omitempty`
}

type UsersStatus struct {
	AuthyId     int      `json:"authy_id"`
	Confirmed   bool     `json:"confirmed"`
	Registered  bool     `json:"registered"`
	CountryCode int      `json:"country_code"`
	PhoneNumber string   `json:"phone_number"`
	Devices     []string `json:"devices"`
}

type UserStatusReq struct {
	AuthyId int `json:"authy_id"`
}

type UserStatusResp struct {
	Status  *UsersStatus `json:"status"`
	Message string       `json:"message",omitempty`
	Success bool         `json:"success",omitempty`
}

type UserRemoveReq struct {
	AuthyId int `json:"authy_id"`
}

type UserRemoveResp struct {
	Message string `json:"message",omitempty`
	Success bool   `json:"success",omitempty`
}

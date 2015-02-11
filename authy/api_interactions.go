package authy

// NewUserReq used to make NewUser API call
type NewUserReq struct {
	User *User `json:"user"`
}

// NewUserResp for JSON new user response
type NewUserResp struct {
	User    *User             `json:"user"`
	Errors  map[string]string `json:"errors",omitempty`
	Message string            `json:"message",omitempty`
	Success bool              `json:"success",omitempty`
}

// VerifyReq used to make token verification API call
type VerifyReq struct {
	Token        string `json:"token"`
	AuthyID      int    `json:"authy_id"`
	Force        bool   `json:"-"`
	CustomAction string `json:"-"`
}

// VerifyResp for JSON token verification response
type VerifyResp struct {
	Token   string            `json:"token"`
	Errors  map[string]string `json:"errors",omitempty`
	Message string            `json:"message",omitempty`
	Success bool              `json:"success",omitempty`
}

// SMSReq used to make SMS token API call
type SMSReq struct {
	AuthyID      int    `json:"authy_id"`
	Force        bool   `json:"-"`
	Shortcode    bool   `json:"-"`
	CustomAction string `json:"-"`
}

// SMSResp for JSON SMS token response
type SMSResp struct {
	Errors    map[string]string `json:"errors",omitempty`
	Message   string            `json:"message",omitempty`
	Cellphone string            `json:"cellphone",omitempty`
	Success   bool              `json:"success",omitempty`
}

// CallReq for phone call token API call
type CallReq struct {
	AuthyID int  `json:"authy_id"`
	Force   bool `json:"-"`
}

// CallResp for JSON call token response
type CallResp struct {
	Errors  map[string]string `json:"errors",omitempty`
	Message string            `json:"message",omitempty`
	Success bool              `json:"success",omitempty`
}

// UsersStatus nested JSON user status response
type UsersStatus struct {
	AuthyID     int      `json:"authy_id"`
	Confirmed   bool     `json:"confirmed"`
	Registered  bool     `json:"registered"`
	CountryCode int      `json:"country_code"`
	PhoneNumber string   `json:"phone_number"`
	Devices     []string `json:"devices"`
}

// UserStatusReq for user status API call
type UserStatusReq struct {
	AuthyID int `json:"authy_id"`
}

// UserStatusResp for JSON user status response
type UserStatusResp struct {
	Status  *UsersStatus `json:"status"`
	Message string       `json:"message",omitempty`
	Success bool         `json:"success",omitempty`
}

// UserRemoveReq for delete user API call
type UserRemoveReq struct {
	AuthyID int `json:"authy_id"`
}

// UserRemoveResp for JSON delete user response
type UserRemoveResp struct {
	Message string `json:"message",omitempty`
	Success bool   `json:"success",omitempty`
}

// AppDetails for App details JSON response
type AppDetails struct {
	Name       string `json:"name"`
	Plan       string `json:"plan"`
	SmsEnabled bool   `json:"sms_enabled"`
	WhiteLabel bool   `json:"white_label"`
}

// AppDetailsResp for JSON app details response
type AppDetailsResp struct {
	App     *AppDetails `json:"app"`
	Message string      `json:"message",omitempty`
	Success bool        `json:"success",omitempty`
}

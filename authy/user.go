package authy

type AuthyUser struct {
	Id          int    `json:"id",omitempty`
	Email       string `json:"email"`
	Cellphone   string `json:"cellphone"`
	CountryCode int    `json:"country_code"`
}

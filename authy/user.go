package authy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthyUser struct {
	Id          int    `json:"id",omitempty`
	Email       string `json:"email"`
	Cellphone   string `json:"cellphone"`
	CountryCode int    `json:"country_code"`
}

func (c *Client) UserStatus(req *UserStatusReq) UserStatusResp {
	path := fmt.Sprintf("/protected/json/users/%+v/status", req.AuthyId)
	apiEndpoint := c.endpoint(path)
	resp, _ := http.Get(apiEndpoint)

	// Unmarshal JSON response
	var verify UserStatusResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &verify)
	return verify
}

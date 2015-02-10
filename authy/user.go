package authy

import (
	"bytes"
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

func (c *Client) RemoveUser(req *UserRemoveReq) UserRemoveResp {
	path := fmt.Sprintf("/protected/json/users/%+v/delete", req.AuthyId)
	apiEndpoint := c.endpoint(path)
	resp, _ := http.Post(apiEndpoint, "application/json", bytes.NewReader([]byte{}))

	// Unmarshal JSON response
	var userResp UserRemoveResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &userResp)
	return userResp
}

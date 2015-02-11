package authy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User is the user created via the API
type User struct {
	ID          int    `json:"id",omitempty`
	Email       string `json:"email"`
	Cellphone   string `json:"cellphone"`
	CountryCode int    `json:"country_code"`
}

// UserStatus gets status of an existing user
// http://docs.authy.com/#section-User_Status
func (c *Client) UserStatus(req *UserStatusReq) UserStatusResp {
	path := fmt.Sprintf("/protected/json/users/%+v/status", req.AuthyID)
	apiEndpoint := c.endpoint(path)
	resp, _ := http.Get(apiEndpoint)

	// Unmarshal JSON response
	var status UserStatusResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &status)
	return status
}

// RemoveUser will delete a user from your account
// http://docs.authy.com/#section-Deleting_User
func (c *Client) RemoveUser(req *UserRemoveReq) UserRemoveResp {
	path := fmt.Sprintf("/protected/json/users/%+v/delete", req.AuthyID)
	apiEndpoint := c.endpoint(path)
	resp, _ := http.Post(apiEndpoint, "application/json", bytes.NewReader([]byte{}))

	// Unmarshal JSON response
	var userResp UserRemoveResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &userResp)
	return userResp
}

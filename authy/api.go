package authy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	SANDBOX_URL = "http://sandbox-api.authy.com"
	API_URL     = "http://api.authy.com"
	REQ_FORMAT  = "json"
)

type Client struct {
	ApiKey string
	apiURL string
}

func NewClient(key string) *Client {
	return &Client{
		ApiKey: key,
		apiURL: ApiURL(false),
	}
}

func NewSandboxClient(key string) *Client {
	return &Client{
		ApiKey: key,
		apiURL: ApiURL(true),
	}
}

func ApiURL(sandbox bool) string {
	if sandbox {
		return SANDBOX_URL
	}
	return API_URL
}

func (c *Client) endpoint(path string) string {
	return c.apiURL + path + "?api_key=" + c.ApiKey
}

func (c *Client) NewUser(req *NewUserReq) NewUserResp {
	// Marshal NewUserReq into JSON to POST
	body, _ := json.Marshal(req)
	resp, _ := http.Post(c.endpoint("/protected/json/users/new"), "application/json", bytes.NewReader(body))

	// Unmarshal JSON response
	var userResp NewUserResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &userResp)
	return userResp
}

func (c *Client) Verify(req *VerifyReq) VerifyResp {
	path := fmt.Sprintf("/protected/json/verify/%+v/%+v", req.Token, req.AuthyId)
	resp, _ := http.Get(c.endpoint(path))

	// Unmarshal JSON response
	var verify VerifyResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &verify)
	return verify
}

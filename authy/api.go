package authy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// SandboxURL is the sandbox test API endpoint
	SandboxURL = "http://sandbox-api.authy.com"
	// APIURL is the production AI endpoint
	APIURL = "http://api.authy.com"
)

// Client for making API calls
// Takes an APIKey
type Client struct {
	APIKey string
	apiURL string
}

// NewClient takes an APIKey and returns a production API Client
func NewClient(key string) *Client {
	return &Client{
		APIKey: key,
		apiURL: APIEndpoint(false),
	}
}

// NewSandboxClient takes an APIKey and returns a sandbox API Client
func NewSandboxClient(key string) *Client {
	return &Client{
		APIKey: key,
		apiURL: APIEndpoint(true),
	}
}

// APIEndpoint returns the URL for the API
func APIEndpoint(sandbox bool) string {
	if sandbox {
		return SandboxURL
	}
	return APIURL
}

func (c *Client) endpoint(path string) string {
	return c.apiURL + path + "?api_key=" + c.APIKey
}

// NewUser creates a new Authy User
// http://docs.authy.com/#section-Enabling_two-factor_on_a_user
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

// Verify will verify a token for an Authy User
// http://docs.authy.com/#section-Verifying_a_token
func (c *Client) Verify(req *VerifyReq) VerifyResp {
	path := fmt.Sprintf("/protected/json/verify/%+v/%+v", req.Token, req.AuthyID)
	apiEndpoint := c.endpoint(path)
	if req.Force {
		apiEndpoint += "&force=true"
	}
	if req.CustomAction != "" {
		apiEndpoint += fmt.Sprintf("&action=%+v", req.CustomAction)
	}
	resp, _ := http.Get(apiEndpoint)

	// Unmarshal JSON response
	var verify VerifyResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &verify)
	return verify
}

package authy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SMS will send a token via SMS
// http://docs.authy.com/#section-SMS_Tokens.
func (c *Client) SMS(req *SMSReq) SMSResp {
	path := fmt.Sprintf("/protected/json/sms/%+v", req.AuthyID)
	apiEndpoint := c.endpoint(path)
	if req.Force {
		apiEndpoint += "&force=true"
	}
	if req.Shortcode {
		apiEndpoint += "&shortcode=true"
	}
	if req.CustomAction != "" {
		apiEndpoint += fmt.Sprintf("&action=%+v", req.CustomAction)
	}
	resp, _ := http.Get(apiEndpoint)

	// Unmarshal JSON response
	var verify SMSResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &verify)
	return verify
}

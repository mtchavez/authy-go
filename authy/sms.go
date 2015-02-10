package authy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) SMS(req *SMSReq) SMSResp {
	path := fmt.Sprintf("/protected/json/sms/%+v", req.AuthyId)
	apiEndpoint := c.endpoint(path)
	if req.Force {
		apiEndpoint += "&force=true"
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

package authy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Call will send a toke via phone call
// http://docs.authy.com/#section-Phone_Call_Tokens.
func (c *Client) Call(req *CallReq) CallResp {
	path := fmt.Sprintf("/protected/json/call/%+v", req.AuthyID)
	apiEndpoint := c.endpoint(path)
	if req.Force {
		apiEndpoint += "&force=true"
	}
	resp, _ := http.Get(apiEndpoint)

	// Unmarshal JSON response
	var verify CallResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &verify)
	return verify
}

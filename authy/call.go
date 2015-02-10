package authy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) Call(req *CallReq) CallResp {
	path := fmt.Sprintf("/protected/json/call/%+v", req.AuthyId)
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

package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// AppDetails API call to get details on your Authy App
func (c *Client) AppDetails() AppDetailsResp {
	apiEndpoint := c.endpoint("/protected/json/app/details")
	resp, _ := http.Get(apiEndpoint)

	// Unmarshal JSON response
	var details AppDetailsResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &details)
	return details
}

func (c *Client) AppStats() AppStatsResp {
	apiEndpoint := c.endpoint("/protected/json/app/stats")
	resp, _ := http.Get(apiEndpoint)

	// Unmarshal JSON response
	var stats AppStatsResp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &stats)
	return stats
}

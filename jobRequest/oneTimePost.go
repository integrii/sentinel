package jobRequest

import (
	"net/http"
	"net/url"
)

// OneTimePOST sends a single POST request that schedules a call to the specified
// URL with the specified post parameters at the specified time
func OneTimePOST(sentinelURL string, jobRequest JobRequest) (*http.Response, error) {
	var params url.Values
	for name, value := range jobRequest.Parameters {
		params.Add(name, value)
	}
	return http.PostForm(sentinelURL, params)
}

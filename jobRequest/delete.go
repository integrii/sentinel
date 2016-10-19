package jobRequest

import (
	"net/http"
	"strings"
)

// Delete deletes the specified job id
func Delete(sentinelURL string, jobID string) (*http.Response, error) {

	// convert bytes to a byte reader
	body := strings.NewReader(jobID)

	// do request and return response and error to caller
	return http.Post(sentinelURL, "text/plain", body)
}

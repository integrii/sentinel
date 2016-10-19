package jobRequest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// SendJobRequest sends a single job request to the sentinel server specified
func SendJobRequest(sentinelURL string, jobRequest JobRequest) (*http.Response, error) {

	// encode the job request into a slice of bytes
	b, err := json.Marshal(jobRequest)
	if err != nil {
		log.Println("Error sending OneTimePOST", err)
	}

	log.Println("Sending POST body:", string(b))

	// convert bytes to a byte reader
	body := bytes.NewReader(b)

	// do request and return response and error to caller
	return http.Post(sentinelURL, "text/plain", body)
}

package jobRequest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// OneTimePOST sends a single POST request that schedules a call to the specified
// URL with the specified post parameters at the specified time
func OneTimePOST(sentinelURL string, jobRequest JobRequest) (*http.Response, error) {

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

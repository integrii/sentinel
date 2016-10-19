package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/integrii/sentinel/jobRequest"
	"github.com/integrii/sentinel/jobTypes"
	"github.com/integrii/sentinel/schedule"
)

const testListenPort = "9010"

var testPostbacksRecieved int

var testServerURLRoot string
var testServerURLDelete string

// Start a sentinel webserver and a test postback endpoint, then
// run tests that make jobs and check that they triggered.
func TestMain(m *testing.M) {
	// test on a random port
	listenPort = 9009
	listenIP = "127.0.0.1" // avoids firewall prompts
	testServerURLRoot = "http://127.0.0.1:" + strconv.Itoa(listenPort)
	testServerURLDelete = "http://127.0.0.1:" + strconv.Itoa(listenPort) + "/delete"

	// start sentinel test postback web server
	go startTestServer()

	// start main as if its normal operation
	go main()

	// wait one second for listeners to come up
	<-time.After(1 * time.Second)

	// run tests and exit
	os.Exit(m.Run())
}

// startTestServer starts a test http service
func startTestServer() {
	log.Println("Testing server listening on port", testListenPort)
	testServer := http.NewServeMux()
	testServer.HandleFunc("/", testHandler)
	err := http.ListenAndServe("127.0.0.1:"+testListenPort, testServer)
	if err != nil {
		log.Println("Testing HTTP Server error:", err)
	}
}

// testHandler handles postbacks that confirm a test ran
func testHandler(w http.ResponseWriter, req *http.Request) {
	testPostbacksRecieved++
	log.Println("Test handler got a postback! #", testPostbacksRecieved)
}

// TestScheduleJob tests the scheduling and running of a job
func TestScheduleJob(t *testing.T) {
	testPostbackRecievedStart := testPostbacksRecieved

	jobID, err := makeTestJob()
	if err != nil {
		t.Error("Error scheduling test job:", err)
		t.FailNow()
	}
	log.Println("Made test job with ID", jobID)

	// wait two seconds and see if a postback came in
	ticker := time.NewTicker(time.Second * 2)
	<-ticker.C
	if !(testPostbacksRecieved == testPostbackRecievedStart+1) {
		t.Fail()
	}
}

// TestDeleteJob creates a job then deletes it
func TestDeleteJob(t *testing.T) {
	jobID, err := makeTestJob()
	if err != nil {
		t.Error("Error scheduling test job:", err)
		t.FailNow()
	}
	log.Println("Made test job with ID", jobID)

	// delete the job we just made
	resp, err := jobRequest.Delete(testServerURLDelete, jobID)
	if err != nil {
		t.Error("Error when deleting job from server", err)
		t.FailNow()
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Error when deleting job from server. Server returned non-200")
		t.FailNow()
	}

}

// makes a test job and returns the ID of it
func makeTestJob() (string, error) {
	var jobID string
	var err error

	// make a new test job
	params := make(map[string]string)
	params["test"] = "true"
	s := schedule.NewOneTimeSchedule(time.Now().Add(time.Second).Unix())
	jr := jobRequest.New("http://localhost:"+testListenPort, params, s)
	jr.JobType = jobTypes.HTTPPOST

	// Send a job to the server
	log.Println("Sending sentinel scheduling POST.")
	resp, err := jobRequest.SendJobRequest(testServerURLRoot, jr)
	if err != nil {
		fmt.Println("Error when sending one time post to schedule job:", err)
		return jobID, err
	}
	log.Println("Sentinel scheduling response code:", resp.StatusCode)
	responseBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("Sentinel scheduling response body:", string(responseBody))

	jobID = string(responseBody)
	return jobID, err
}

package main

import (
	"fmt"
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

// Start a sentinel webserver and a test postback endpoint, then
// run tests that make jobs and check that they triggered.
func TestMain(m *testing.M) {
	// test on a random port
	listenPort = 9009

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

	// make a new test job
	params := make(map[string]string)
	params["test"] = "true"
	s := schedule.NewOneTimeSchedule(time.Now().Add(time.Second))
	jr := jobRequest.New("http://localhost:"+testListenPort, params, s)
	jr.JobType = jobTypes.HTTPPOSTJSON

	// Send a job to the server
	log.Println("Sending sentinel scheduling POST.")
	resp, err := jobRequest.SendJobRequest("http://127.0.0.1:"+strconv.Itoa(listenPort), jr)
	if err != nil {
		fmt.Println("Error when sending one time post to schedule job:", err)
		t.FailNow()
	}
	log.Println("Sentinel scheduling response code:", resp.StatusCode)
	// responseBody, _ := ioutil.ReadAll(resp.Body)
	// log.Println("Sentinel scheduling response body:", string(responseBody))

	// wait two seconds and see if a postback came in
	ticker := time.NewTicker(time.Second * 2)
	<-ticker.C
	if !(testPostbacksRecieved == testPostbackRecievedStart+1) {
		t.Fail()
	}
}

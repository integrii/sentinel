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

	// run tests and exit
	os.Exit(m.Run())
}

// startTestServer starts a test http service
func startTestServer() {
	log.Println("Testing server listening on port", testListenPort)
	testServer := http.NewServeMux()
	testServer.HandleFunc("/", testHandler)
	err := http.ListenAndServe(":"+testListenPort, testServer)
	if err != nil {
		log.Println("Testing HTTP Server error:", err)
	}
}

// testHandler handles postbacks that confirm a test ran
func testHandler(w http.ResponseWriter, req *http.Request) {
	testPostbacksRecieved++
}

// TestScheduleJob tests the scheduling and running of a job
func TestScheduleJob(t *testing.T) {
	testPostbackRecievedStart := testPostbacksRecieved

	// make a new test job
	params := make(map[string]string)
	params["test"] = "true"
	s := schedule.NewOneTimeSchedule(time.Now())
	jr := jobRequest.New(params, s)

	// Send a job to the server
	_, err := jobRequest.OneTimePOST("http://127.0.0.1:"+strconv.Itoa(listenPort), jr)
	if err != nil {
		fmt.Println("Error when sending one time post to schedule job:", err)
		t.FailNow()
	}

	// wait two seconds and see if a postback came in
	ticker := time.NewTicker(time.Second * 2)
	<-ticker.C
	if !(testPostbacksRecieved > testPostbackRecievedStart+1) {
		t.Fail()
	}
}

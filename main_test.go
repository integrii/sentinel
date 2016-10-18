package main

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/integrii/sentinel/jobRequest"
	"github.com/integrii/sentinel/schedule"
)

const testListenPort = 9010

var testPostbacksRecieved int

// Start a sentinel webserver
func TestMain(m *testing.M) {
	// test on a random port
	listenPort = 9009

	// start main as if its normal operation
	main()

	// start sentinel test postback web server
	go startTestServer()

	// run tests and exit
	os.Exit(m.Run())
}

func startTestServer() {
	http.HandleFunc("/", testHandler)
	http.ListenAndServe(":"+string(testListenPort), nil)
}

// testHandler handles postbacks that confirm a test ran
func testHandler(w http.ResponseWriter, req *http.Request) {
	testPostbacksRecieved++
}

func TestScheduleJob(t *testing.T) {
	testPostbackRecievedStart := testPostbacksRecieved

	// make a new test job
	params := make(map[string]string)
	params["test"] = "true"
	s := schedule.NewOneTimeSchedule(time.Now())
	jr := jobRequest.New(params, s)

	// Send a job to the server
	jobRequest.OneTimePOST("localhost:"+string(listenPort), jr)

	// wait two seconds and see if a postback came in
	ticker := time.NewTicker(time.Second * 2)
	<-ticker.C
	if !(testPostbacksRecieved > testPostbackRecievedStart+1) {
		t.Fail()
	}
}

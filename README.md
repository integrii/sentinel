# Sentinel
Schedule remote web calls with an API call to this service.  Kind of like cron, but a web service.  Not backed by any DB.  Currently wipes all cron jobs on each start.

#### Installing

`go get github.com/integrii/sentinel`

Then, start the server by running `sentinel`.  It listens on port 80.


## Scheduling a task

Send a POST to the root of the server with the task JSON in the body.  If you are using a go application, the `jobRequest` sub-package here contains helpful functions for creating and sending these requests.  The main_test.go file also contains a helpful example.

A web request for a simple POST job that sends a post header for each property looks like this:

```
{
  "JobType": "HTTPPOSTJSON",
  "ServerURL": "https://test.com/test",
  "Parameters": {
    "test": "true",
    "great": "Absolutely!"
  },
  "Schedule": {
    "Time": 1476853274
  }
}
```

The server will return a 200 status header along with an ID for the scheduled task in the body in plain text.

#### The possible `JobType` options are:
- **HTTPPOST**: Causes a post where the `Parameters` section in the request is sent as POST form headers.

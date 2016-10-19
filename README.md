# Sentinel

Schedule remote web calls with an API call to this service.  Kind of like cron, but a web service.  Not backed by any DB.  Currently wipes all cron jobs on each start.

#### Installing

`go get github.com/integrii/sentinel`

Then, start the server by running `sentinel`.  It listens on port 80.


## Scheduling a Task

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


## Deleting a Task

Send a POST or GET to the `/test` endpoint of the server with the ID of the task that you want to delete in the body as plain text.  You can expect a 200 from this or a `401` if the job did not exist.


## Insomnia Workspace

A insomnia workspace is included that can generate Sentinel web request examples in nearly any language you desire.  Simply download [Insomnia](http://insomnia.rest) and load the `Insomnia-Workspace.json` file.

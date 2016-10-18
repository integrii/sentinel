package main

// JobType defines a specific job string identifier.
type JobType string

const JOB_HTTPGET = JobType("HTTPGET")
const JOB_HTTPPOST = JobType("HTTPPOST")
const JOB_HTTPPOSTJSON = JobType("HTTPPOSTJSON")

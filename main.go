package main

import (
	"fmt"
	"log"
	"my-protobuf/jobsearch"
	"time"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// basic.BasicHello()
	// basic.BasicUser()
	// basic.ProtoToJsonUser()
	// basic.BasicUserGroup()
	jobsearch.JobSearchSoftware()
	jobsearch.JobSearchCandidate()
}

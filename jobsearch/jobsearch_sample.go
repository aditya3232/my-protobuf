package jobsearch

import (
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/dummy"
	"my-protobuf/protogen/jobsearch"

	"log"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 666,
		Application: &basic.Application{
			Version: "1.0.0",
			Name:    "ProtoBuf JobSearch Software",
			Platforms: []string{
				"Mac",
				"Windows",
				"Linux",
			},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("Software: ", string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 777,
		Application: &dummy.Application{
			ApplicationId:     888,
			ApplicantFullName: "John Doe",
			Phone:             "123-456-7890",
			Email:             "shazam@yahoo.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Println("Candidate: ", string(jsonBytes))
}

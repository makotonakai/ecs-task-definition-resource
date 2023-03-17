package main

import (
	"os"
	"fmt"
	"encoding/json"

	"github.com/makotonakai/ecs-task-definition-resource/resource"
)

type Request struct {
	Params  resource.Params  `json:"params"`
}

type Response struct {
	Version resource.Version `json:"version"`
}

func main() {

	var request Request
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&request)
	if err != nil {
			fmt.Fprintf(os.Stderr, "failed to decode: %s\n", err.Error())
			os.Exit(1)
			return
	}

	params := request.Params
	fmt.Fprintf(os.Stderr, "AWS Access Key ID: %v\n", params.AccessKeyId)
	fmt.Fprintf(os.Stderr, "AWS Secret Access Key: %v\n", params.SecretAccessKey)
	fmt.Fprintf(os.Stderr, "AWS Region: %v\n", params.Region)
	fmt.Fprintf(os.Stderr, "Task Definition: %v\n", params.TaskDefinition)

	response := Response{}
	response.Version = resource.Version{Ref: "static"}

	json.NewEncoder(os.Stdout).Encode(response)
}


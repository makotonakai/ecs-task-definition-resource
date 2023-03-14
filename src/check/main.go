package main

import (
	"os"
	"fmt"
	"time"
	"encoding/json"
	"github.com/makotonakai/ecs-task-definition-resource/resource"
)

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
}

type Response []resource.Version

func main() {
	
	var request Request
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&request)
	if err != nil {
			fmt.Fprintf(os.Stderr, "failed to decode: %s\n", err.Error())
			os.Exit(1)
			return
	}

	fmt.Fprintf(os.Stderr, "source: %v\n", request.Source)

	response := Response{}
	response = append(response, resource.Version{Date: time.Now().String()})

	json.NewEncoder(os.Stdout).Encode(response)
}


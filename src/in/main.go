package main

import (
	"os"
	"encoding/json"
	"github.com/makotonakai/ecs-task-definition-resource/resource"
)

type Response struct {
	Version resource.Version `json:"version"`
}

func main() {

	response := Response{}

	version := resource.Version{Ref: ""}
	response.Version = version

	json.NewEncoder(os.Stdout).Encode(response)
	
}


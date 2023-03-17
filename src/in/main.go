package main

import (
	"os"
	"fmt"
	"time"
	"encoding/gob"
	"encoding/json"
	"github.com/makotonakai/ecs-task-definition-resource/resource"
)

type Response struct {
	Version resource.Version `json:"version"`
	MetaData []resource.Metadata `json:"metadata"`
}

func main() {

	response := Response{}

	version := resource.Version{Ref: "static"}
	response.Version = version

	metadata := resource.Metadata{Name: "key", Value: "value"}
	response.MetaData = append(response.MetaData, metadata)

	json.NewEncoder(os.Stdout).Encode(response)
	
}


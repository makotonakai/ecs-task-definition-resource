package main

import (
	"os"
	"encoding/json"
	"github.com/makotonakai/ecs-task-definition-resource/resource"
)


type Response struct {
	Version []resource.Version `json:"version"`
}

func main() {
	
	response := Response{}
	json.NewEncoder(os.Stdout).Encode(response)

}


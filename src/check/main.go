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
	response.Version = append(response.Version, resource.Version{Ref: ""})
	json.NewEncoder(os.Stdout).Encode(response)

}


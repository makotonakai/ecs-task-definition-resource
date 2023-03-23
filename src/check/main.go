package main

import (
	"os"
	"log"
	"encoding/json"
	"github.com/makotonakai/ecs-task-definition-resource/resource"
)


type Response struct {
	Version []resource.Version `json:"version"`
}

func main() {
	
	response := Response{}
	err := json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}

}


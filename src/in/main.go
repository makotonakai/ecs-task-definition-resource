package main

import (
	"os"
	"log"
	"encoding/json"
	"github.com/makotonakai/ecs-task-definition-resource/resource"
)

// type Request struct {
// 	Version resource.Version `json:"version"`
// }


type Response struct {
	Version resource.Version `json:"version"`
}

func main() {

	// var request Request
	// decoder := json.NewDecoder(os.Stdin)
	// err := decoder.Decode(&request)
	// if err != nil {
	// 		fmt.Fprintf(os.Stderr, "failed to decode: %s\n", err.Error())
	// 		os.Exit(1)
	// 		return
	// }

	response := Response{}
	response.Version = resource.Version{Ref: "static"}

	err := json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
	
}


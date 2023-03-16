package main

import (
	"os"
	"fmt"
	"time"
	"encoding/json"

	"github.com/makotonakai/ecs-task-definition-resource/resource"
	"github.com/makotonakai/ecs-task-definition-resource/src/common"
)

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
}

type Response struct {
	Version resource.Version `json:"version"`
	MetaData []resource.Metadata `json:"metadata"`
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

	dest := os.Args[0]

	awsAccessKeyIdFile := fmt.Sprintf("%s/aws_access_key_id", dest)
	awsSecretAccessKeyFile := fmt.Sprintf("%s/aws_secret_access_key", dest)
	awsRegionFile := fmt.Sprintf("%s/aws_region", dest)
	taskDefinisionFile := fmt.Sprintf("%s/task_definition_json", dest)

	// task_definition.jsonの構造体をバイト列に変換してファイルに書き込む
	awsAccessKeyId := common.GetStringFromFile(awsAccessKeyIdFile)
	awsSecretAccessKey := common.GetStringFromFile(awsSecretAccessKeyFile)
	awsRegion := common.GetStringFromFile(awsRegionFile)
	taskDefinision := common.GetTaskDefinitionFromFile(taskDefinisionFile)

	fmt.Fprintf(os.Stderr, "AWS Access Key ID: %v\n", awsAccessKeyId)
	fmt.Fprintf(os.Stderr, "AWS Secret Access Key: %v\n", awsSecretAccessKey)
	fmt.Fprintf(os.Stderr, "AWS Region: %v\n", awsRegion)
	fmt.Fprintf(os.Stderr, "Task Definition: %v\n", taskDefinision)

	response := Response{}
	response.Version = resource.Version{Date: time.Now().String()}
	response.MetaData = []resource.Metadata{}

	json.NewEncoder(os.Stdout).Encode(response)
}


package main

import (
	"os"
	"fmt"
	"time"
	"encoding/gob"
	"encoding/json"
	"github.com/makotonakai/ecs-task-definition-resource/resource"
)

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
}

type Response struct {
	Version []resource.Version `json:"version"`
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

	awsAccessKeyIdPath := fmt.Sprintf("%s/aws_access_key_id", dest)
	awsSecretAccessKeyPath := fmt.Sprintf("%s/aws_secret_access_key", dest)
	awsRegionPath := fmt.Sprintf("%s/aws_region", dest)
	taskDefinisionJSONPath := fmt.Sprintf("%s/task_definition_json", dest)

	awsAccessKeyIdFile, _ := os.Create(awsAccessKeyIdPath)
	awsSecretAccessKeyFile, _ := os.Create(awsSecretAccessKeyPath)
	awsRegionFile, _ := os.Create(awsRegionPath)
	taskDefinisionFile, _ := os.Create(taskDefinisionJSONPath)

	defer awsAccessKeyIdFile.Close()
	defer awsSecretAccessKeyFile.Close()
	defer awsRegionFile.Close()
	defer taskDefinisionFile.Close()

	source := request.Source

	awsAccessKeyIdFile.WriteString(source.AccessKeyId)
	awsSecretAccessKeyFile.WriteString(source.SecretAccessKey)
	awsRegionFile.WriteString(source.Region)

	// task_definition.jsonの構造体をバイト列に変換してファイルに書き込む
	encoder := gob.NewEncoder(taskDefinisionFile)
	encoder.Encode(source.TaskDefinition)

	response := Response{}

	version := resource.Version{Date: time.Now().String()}
	response.Version = append(response.Version, version)

	metadata := resource.Metadata{Name: "key", Value: "value"}
	response.MetaData = append(response.MetaData, metadata)

	json.NewEncoder(os.Stdout).Encode(response)
}


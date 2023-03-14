package main

import (
	"os"
	"fmt"
	"time"
	"encoding/gob"
	"encoding/json"
	"github.com/makotonakai/ecs-task-definition-resource/model"
	"github.com/makotonakai/ecs-task-definition-resource/resource"
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

	// task_definition.jsonの構造体をバイト列に変換してファイルに書き込む
	awsAccessKeyId := GetStringFromFile(awsAccessKeyIdFile)
	awsSecretAccessKey := GetStringFromFile(awsSecretAccessKeyFile)
	awsRegion := GetStringFromFile(awsRegionFile)

	fmt.Fprintf(os.Stderr, "AWS Access Key ID: %v\n", awsAccessKeyId)
	fmt.Fprintf(os.Stderr, "AWS Secret Access Key: %v\n", awsSecretAccessKey)
	fmt.Fprintf(os.Stderr, "AWS Region: %v\n", awsRegion)

	response := Response{}
	response.Version = resource.Version{Date: time.Now().String()}
	response.MetaData = []resource.Metadata{}

	json.NewEncoder(os.Stdout).Encode(response)
}

func GetStringFromFile(fileName string) string {
	fp, err := os.Open(fileName)
	if err != nil {
			panic(err)
	}
	defer fp.Close()

	buf := make([]byte, 1024)
	for {
			n, err := fp.Read(buf)
			if n == 0 {
					break
			}
			if err != nil {
					panic(err)
			}
	}
	return string(buf)
}

func GetTaskDefinitionFromFile(fileName string) model.TaskDefinision {
	fp, err := os.Open(fileName)
	if err != nil {
			panic(err)
	}
	defer fp.Close()
	
	var taskDefinition model.TaskDefinision
	decoder := gob.NewDecoder(fp)
	decoder.Decode(&taskDefinition)
	return taskDefinition
}


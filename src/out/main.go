package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/makotonakai/ecs-task-definition-resource/resource"
)

type Request struct {
	Params  resource.Params  `json:"params"`
}

type Response struct {
	Version resource.Version `json:"version"`
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

	params := request.Params
	awsAccessKeyId := params.AccessKeyId
	awsSecretAccessKey := params.SecretAccessKey
	awsRegion := params.Region
	taskDefinition := params.TaskDefinition

	fmt.Fprintf(os.Stderr, "AWS Access Key ID: %v\n", awsAccessKeyId)
	fmt.Fprintf(os.Stderr, "AWS Secret Access Key: %v\n", awsSecretAccessKey)
	fmt.Fprintf(os.Stderr, "AWS Region: %v\n", awsRegion)
	fmt.Fprintf(os.Stderr, "Task Definition: %v\n", taskDefinition)

	cred := credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, "")
	sess := session.Must(session.NewSession())

	svc := ecs.New(
		sess,
		aws.NewConfig().WithRegion(awsRegion).WithCredentials(cred),
	)
	input := &ecs.RegisterTaskDefinitionInput{
			ContainerDefinitions: []*ecs.ContainerDefinition{
					{
							Command: []*string{
									aws.String("sleep"),
									aws.String("360"),
							},
							Cpu:       aws.Int64(10),
							Essential: aws.Bool(true),
							Image:     aws.String("busybox"),
							Memory:    aws.Int64(10),
							Name:      aws.String("sleep"),
					},
			},
			Family:      aws.String("sleep360"),
			TaskRoleArn: aws.String(""),
	}

	result, err := svc.RegisterTaskDefinition(input)
	if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case ecs.ErrCodeServerException:
							fmt.Println(ecs.ErrCodeServerException, aerr.Error())
					case ecs.ErrCodeClientException:
							fmt.Println(ecs.ErrCodeClientException, aerr.Error())
					case ecs.ErrCodeInvalidParameterException:
							fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
					default:
							fmt.Println(aerr.Error())
					}
			} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
			}
			return
	}

	fmt.Println(result)

	response := Response{}
	response.Version = resource.Version{Ref: time.Now().String()}

	err = json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}


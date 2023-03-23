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
							Cpu:       aws.Int64(taskDefinition.ContainerDefinitions[0].Cpu),
							Essential: aws.Bool(taskDefinition.ContainerDefinitions[0].Essential),
							Image:     aws.String(taskDefinition.ContainerDefinitions[0].Image),
							Name:      aws.String(taskDefinition.ContainerDefinitions[0].Name),
							HealthCheck: &ecs.HealthCheck{
								Command: []*string{
									aws.String("CMD-SHELL"),
									aws.String("curl -f http://localhost:1991/ping || exit 1"),
								},
								Interval: aws.Int64(taskDefinition.ContainerDefinitions[0].HealthCheck.Interval),
								Retries: aws.Int64(taskDefinition.ContainerDefinitions[0].HealthCheck.Retries),
								StartPeriod: aws.Int64(taskDefinition.ContainerDefinitions[0].HealthCheck.StartPeriod),
								Timeout: aws.Int64(taskDefinition.ContainerDefinitions[0].HealthCheck.Timeout),
							},
							LogConfiguration: &ecs.LogConfiguration{
								LogDriver: aws.String(taskDefinition.ContainerDefinitions[0].LogConfiguration.LogDriver),
								Options: map[string]*string{
									"awslogs-create-group": aws.String(taskDefinition.ContainerDefinitions[0].LogConfiguration.Options.AWSLogsCreateGroup),
									"awslogs-group": aws.String(taskDefinition.ContainerDefinitions[0].LogConfiguration.Options.AWSLogsGroup),
									"awslogs-region": aws.String(taskDefinition.ContainerDefinitions[0].LogConfiguration.Options.AWSLogsRegion),
									"awslogs-stream-prefix": aws.String(taskDefinition.ContainerDefinitions[0].LogConfiguration.Options.AWSLogsStreamPrefix),
								},
							},
						},	
					},
			Cpu: aws.String(taskDefinition.Cpu),
			EphemeralStorage: &ecs.EphemeralStorage{
				SizeInGiB: aws.Int64(taskDefinition.EphemeralStorage.SizeInGiB),
			},
			ExecutionRoleArn: aws.String(taskDefinition.ExecutionRoleARN),
			Family:      aws.String(taskDefinition.Family),
			Memory: aws.String(taskDefinition.Memory),
			NetworkMode: aws.String(taskDefinition.NetworkMode),
			PlacementConstraints: []*ecs.TaskDefinitionPlacementConstraint{},
			RuntimePlatform: &ecs.RuntimePlatform{
				CpuArchitecture: aws.String(taskDefinition.RuntimePlatform.CpuArchitecture),
				OperatingSystemFamily: aws.String(taskDefinition.RuntimePlatform.OperatingSystemFamily),
			},
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


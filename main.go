package main

import (

	"os"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"

)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Failed to load .env: ", err)
	} 

	awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")

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
}

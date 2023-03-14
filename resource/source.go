package resource

import (
	"github.com/makotonakai/ecs-task-definition-resource/model"
)
type Source struct { 
	AccessKeyId string `json:"aws_access_key_id"`
	SecretAccessKey string `json:"aws_secret_access_key"`
	Region string `json:"aws_region"`
	TaskDefinition model.TaskDefinision `json:"task_definition.json"`
}



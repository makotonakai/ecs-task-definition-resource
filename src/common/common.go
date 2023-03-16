package common

import (
	"os"
	"encoding/gob"
	"github.com/makotonakai/ecs-task-definition-resource/model"
)

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


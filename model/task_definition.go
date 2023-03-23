package model

type TaskDefinision struct {
	TaskDefinitionARN string `json:"taskDefinitionArn"`
	ContainerDefinitions []ContainerDefinition `json:"containerDefinitions"`
	Family string `json:"family"`
	ExecutionRoleARN string `json:"executionRoleArn"`
	NetworkMode string `json:"networkMode"`
	Revision int `json:"revision"`
	Volumes []string `json:"volumes"`
	Status string `json:"status"`
	RequiresAttributes []RequiresAttribute `json:"requiresAttributes"`
	PlacementContraints []string `json:"placementConstraints"`
	Compatibilities []string `json:"compatibilities"`
	RequiredCompatibilities []string `json:"requiresCompatibilities"`
	Cpu string `json:"Cpu"`
	Memory string `json:"memory"`
	EphemeralStorage EphemeralStorageConfig `json:"ephemeralStorageConfig"`
	RuntimePlatform RuntimePlatformConfig `json:"runtimePlatform"`
	RegisteredAt string `json:"registeredAt"`
	RegisteredBy string `json:"registeredBy"`
	Tags []Tag `json:"tags"`
}

type ContainerDefinition struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Cpu int64 `json:"Cpus"`
	PortMappings []PortMapping `json:"portMappings"`
	Essential bool `json:"essential"`
	Environment []string `json:"environment"`
	EnvironmentFiles []string `json:"environmentFiles"`
	MountPoints []string `json:"mountPoints"`
	VolumesFrom []string `json:"volumesFrom"`
	LogConfiguration LogConfig `json:"logConfiguration"`
	HealthCheck HealthCheckConfig `json:"healthCheck"`
}

type PortMapping struct {
	Name string `json:"name"`
	ContainerPort int `json:"containerPort"`
	HostPort int `json:"hostPort"`
	Protocol string `json:"protocol"`
	AppProtocol string `json:"appProtocol"`
}

type LogsOptions struct {
	AWSLogsCreateGroup string `json:"awslogs-create-group"`
	AWSLogsGroup string `json:"awslogs-group"`
	AWSLogsRegion string `json:"awslogs-region"`
	AWSLogsStreamPrefix string `json:"awslogs-stream-prefix"`
}

type LogConfig struct {
	LogDriver string `json:"logDriver"`
	Options LogsOptions `json:"options"`
}

type HealthCheckConfig struct {
	Command []string `json:"command"`
	Interval int64 `json:"interval"`
	Timeout int64 `json:"timeout"`
	Retries int64 `json:"retries"`
	StartPeriod int64 `json:"startPeriod"`
}

type RequiresAttribute struct {
	Name string `json:"name"`
}

type EphemeralStorageConfig struct {
	SizeInGiB int64 `json:"sizeInGiB"`
}

type RuntimePlatformConfig struct {
	CpuArchitecture string `json:"CpuArchitecture"`
	OperatingSystemFamily string `json:"operatingSystemFamily"`
}

type Tag struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

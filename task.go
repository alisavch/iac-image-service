package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewTask creates a task.
func NewTask(ctx *pulumi.Context, name string, role *iam.Role, container pulumi.StringOutput) (*ecs.TaskDefinition, error) {
	taskDefinition, err := ecs.NewTaskDefinition(ctx, name, &ecs.TaskDefinitionArgs{
		Tags:                    pulumi.StringMap{"Name": pulumi.String(name)},
		Family:                  pulumi.String(name),
		Cpu:                     pulumi.String("256"),
		Memory:                  pulumi.String("512"),
		NetworkMode:             pulumi.String("awsvpc"),
		RequiresCompatibilities: pulumi.StringArray{pulumi.String("FARGATE")},
		ExecutionRoleArn:        role.Arn,
		ContainerDefinitions:    container,
	})
	if err != nil {
		return nil, err
	}
	return taskDefinition, nil
}

package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewRole creates new role.
func NewRole(ctx *pulumi.Context) (*iam.Role, error) {
	taskExecRole, err := iam.NewRole(ctx, "pulumi-task-exec-role-as", &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`),
	})
	if err != nil {
		return nil, err
	}

	_, err = iam.NewRolePolicyAttachment(ctx, "pulumi-task-exec-policy-as", &iam.RolePolicyAttachmentArgs{
		Role:      taskExecRole.Name,
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
	})
	if err != nil {
		return nil, err
	}
	return taskExecRole, nil
}

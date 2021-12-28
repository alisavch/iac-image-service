package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewLogGroup creates a cloudwatch.
func NewLogGroup(ctx *pulumi.Context, name string) (*cloudwatch.LogGroup, error) {
	logGroup, err := cloudwatch.NewLogGroup(ctx, name, &cloudwatch.LogGroupArgs{}, nil)
	if err != nil {
		return nil, err
	}
	return logGroup, nil
}

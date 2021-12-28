package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewCluster creates a logical grouping of services or tasks.
func NewCluster(ctx *pulumi.Context, name string) (*ecs.Cluster, error) {
	cluster, err := ecs.NewCluster(ctx, name, nil)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

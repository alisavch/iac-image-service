package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewTargetGroup creates a target group that directs requests to one or more registered targets.
func NewTargetGroup(ctx *pulumi.Context, name string, vpc *ec2.LookupVpcResult) (*lb.TargetGroup, error) {
	targetGroup, err := lb.NewTargetGroup(ctx, name, &lb.TargetGroupArgs{
		Port:       pulumi.Int(8080),
		Protocol:   pulumi.String("HTTP"),
		TargetType: pulumi.String("ip"),
		VpcId:      pulumi.String(vpc.Id),
	})
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

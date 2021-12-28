package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewVpc uses virtual private cloud by default.
func NewVpc(ctx *pulumi.Context) (*ec2.LookupVpcResult, error) {
	defaultVPC := true
	vpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{Default: &defaultVPC})
	if err != nil {
		return nil, err
	}
	return vpc, err
}

package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetSubnet gets a range of IP addresses in VPC.
func GetSubnet(ctx *pulumi.Context, vpc *ec2.LookupVpcResult) (*ec2.GetSubnetIdsResult, error) {
	subnets, err := ec2.GetSubnetIds(ctx, &ec2.GetSubnetIdsArgs{VpcId: vpc.Id})
	if err != nil {
		return nil, err
	}
	return subnets, nil
}

// NewDbSubnet creates subnet group for rds.
func NewDbSubnet(ctx *pulumi.Context, name string, subnets *ec2.GetSubnetIdsResult) (*rds.SubnetGroup, error) {
	dbSubnets, err := rds.NewSubnetGroup(ctx, name, &rds.SubnetGroupArgs{
		SubnetIds: pulumi.ToStringArray(subnets.Ids),
	})
	if err != nil {
		return nil, err
	}
	return dbSubnets, nil
}

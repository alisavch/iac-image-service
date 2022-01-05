package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewLoadBalancer creates a single point of contact for the customer.
func NewLoadBalancer(ctx *pulumi.Context, name string, subnets *ec2.GetSubnetIdsResult, securityGroup *ec2.SecurityGroup) (*lb.LoadBalancer, error) {
	loadBalancer, err := lb.NewLoadBalancer(ctx, name, &lb.LoadBalancerArgs{
		Subnets:        pulumi.ToStringArray(subnets.Ids),
		SecurityGroups: pulumi.StringArray{securityGroup.ID()},
	})
	if err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

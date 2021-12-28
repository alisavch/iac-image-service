package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewRDS creates Postgresql rds.
func NewRDS(ctx *pulumi.Context, name string, conf Config, dbSubnets *rds.SubnetGroup, securityGroup *ec2.SecurityGroup) (*rds.Instance, error) {
	database, err := rds.NewInstance(ctx, name, &rds.InstanceArgs{
		InstanceClass:       pulumi.String("db.t3.micro"),
		Engine:              pulumi.String("postgres"),
		AllocatedStorage:    pulumi.Int(10),
		Name:                pulumi.String("conversion_compression_service"),
		Username:            conf.DbUsername,
		Password:            conf.DbPassword,
		SkipFinalSnapshot:   pulumi.Bool(true),
		DbSubnetGroupName:   dbSubnets.ID().ToStringOutput(),
		PubliclyAccessible:  pulumi.Bool(true),
		VpcSecurityGroupIds: pulumi.StringArray{securityGroup.ID().ToStringOutput()},
	})
	if err != nil {
		return nil, err
	}
	return database, nil
}

package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewSecurityGroup creates virtual firewall to control incoming and outgoing traffic.
func NewSecurityGroup(ctx *pulumi.Context, name string, vpc *ec2.LookupVpcResult) (*ec2.SecurityGroup, error) {
	securityGroup, err := ec2.NewSecurityGroup(ctx, name, &ec2.SecurityGroupArgs{
		VpcId: pulumi.String(vpc.Id),
		Egress: ec2.SecurityGroupEgressArray{
			ec2.SecurityGroupEgressArgs{
				Protocol:       pulumi.String("-1"),
				FromPort:       pulumi.Int(0),
				ToPort:         pulumi.Int(0),
				CidrBlocks:     pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				Ipv6CidrBlocks: pulumi.StringArray{pulumi.String("::/0")},
			},
		},
		Ingress: ec2.SecurityGroupIngressArray{
			ec2.SecurityGroupIngressArgs{
				Protocol:       pulumi.String("tcp"),
				FromPort:       pulumi.Int(22),
				ToPort:         pulumi.Int(22),
				CidrBlocks:     pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				Ipv6CidrBlocks: pulumi.StringArray{pulumi.String("::/0")},
			},
			ec2.SecurityGroupIngressArgs{
				Protocol:       pulumi.String("tcp"),
				FromPort:       pulumi.Int(80),
				ToPort:         pulumi.Int(80),
				CidrBlocks:     pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				Ipv6CidrBlocks: pulumi.StringArray{pulumi.String("::/0")},
			},
			ec2.SecurityGroupIngressArgs{
				Protocol:       pulumi.String("tcp"),
				FromPort:       pulumi.Int(8080),
				ToPort:         pulumi.Int(8080),
				CidrBlocks:     pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				Ipv6CidrBlocks: pulumi.StringArray{pulumi.String("::/0")},
			},
			ec2.SecurityGroupIngressArgs{
				Protocol:       pulumi.String("tcp"),
				FromPort:       pulumi.Int(5671),
				ToPort:         pulumi.Int(5671),
				CidrBlocks:     pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				Ipv6CidrBlocks: pulumi.StringArray{pulumi.String("::/0")},
			},
			ec2.SecurityGroupIngressArgs{
				Protocol:       pulumi.String("tcp"),
				FromPort:       pulumi.Int(5432),
				ToPort:         pulumi.Int(5432),
				CidrBlocks:     pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				Ipv6CidrBlocks: pulumi.StringArray{pulumi.String("::/0")},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return securityGroup, nil
}

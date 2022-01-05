package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API contains a service environment.
type API struct {
	Config
	Cluster       *ecs.Cluster
	Task          *ecs.TaskDefinition
	Subnets       *ec2.GetSubnetIdsResult
	SecurityGroup *ec2.SecurityGroup
	TargetGroup   *lb.TargetGroup
	Listener      *lb.Listener
}

// NewAPI configures API.
func NewAPI(config Config, cluster *ecs.Cluster, task *ecs.TaskDefinition, subnets *ec2.GetSubnetIdsResult, securityGroup *ec2.SecurityGroup, targetGroup *lb.TargetGroup, listener *lb.Listener) API {
	return API{Config: config, Cluster: cluster, Task: task, Subnets: subnets, SecurityGroup: securityGroup, TargetGroup: targetGroup, Listener: listener}
}

// NewAPIService creates a service.
func NewAPIService(ctx *pulumi.Context, name string, service API) error {
	_, err := ecs.NewService(ctx, name, &ecs.ServiceArgs{
		Tags:           pulumi.StringMap{"Name": pulumi.String("ecs-service-api-as")},
		Cluster:        service.Cluster.Arn,
		DesiredCount:   pulumi.Int(1),
		LaunchType:     pulumi.String("FARGATE"),
		TaskDefinition: service.Task.Arn,
		NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
			AssignPublicIp: pulumi.Bool(true),
			Subnets:        pulumi.ToStringArray(service.Subnets.Ids),
			SecurityGroups: pulumi.StringArray{service.SecurityGroup.ID()},
		},
		LoadBalancers: ecs.ServiceLoadBalancerArray{
			ecs.ServiceLoadBalancerArgs{
				TargetGroupArn: service.TargetGroup.Arn,
				ContainerName:  service.Config.APIEnv.Name,
				ContainerPort:  pulumi.Int(8080),
			},
		},
	}, pulumi.DependsOn([]pulumi.Resource{service.Listener}))
	if err != nil {
		return err
	}
	return nil
}

// Consumer contains a service environment.
type Consumer struct {
	Cluster       *ecs.Cluster
	Task          *ecs.TaskDefinition
	Subnets       *ec2.GetSubnetIdsResult
	SecurityGroup *ec2.SecurityGroup
}

// NewConsumer configures consumer.
func NewConsumer(cluster *ecs.Cluster, task *ecs.TaskDefinition, subnets *ec2.GetSubnetIdsResult, securityGroup *ec2.SecurityGroup) Consumer {
	return Consumer{Cluster: cluster, Task: task, Subnets: subnets, SecurityGroup: securityGroup}
}

// NewConsumerService creates a service.
func NewConsumerService(ctx *pulumi.Context, name string, service Consumer) error {
	_, err := ecs.NewService(ctx, name, &ecs.ServiceArgs{
		Tags:           pulumi.StringMap{"Name": pulumi.String(name)},
		Cluster:        service.Cluster.Arn,
		DesiredCount:   pulumi.Int(1),
		LaunchType:     pulumi.String("FARGATE"),
		TaskDefinition: service.Task.Arn,
		NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
			AssignPublicIp: pulumi.Bool(true),
			Subnets:        pulumi.ToStringArray(service.Subnets.Ids),
			SecurityGroups: pulumi.StringArray{service.SecurityGroup.ID().ToStringOutput()},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

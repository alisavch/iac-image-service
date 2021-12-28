package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		config := NewConfig(ctx)

		vpc, err := NewVpc(ctx)
		if err != nil {
			return err
		}

		securityGroup, err := NewSecurityGroup(ctx, "pulumi-sg-as", vpc)
		if err != nil {
			return err
		}

		cluster, err := NewCluster(ctx, "pulumi-cluster-as")
		if err != nil {
			return err
		}

		taskExecRole, err := NewRole(ctx)
		if err != nil {
			return err
		}

		subnets, err := GetSubnet(ctx, vpc)
		if err != nil {
			return err
		}

		logAPI, err := NewLogGroup(ctx, "pulumi-logs-api")
		if err != nil {
			return err
		}

		logConsumer, err := NewLogGroup(ctx, "pulumi-logs-consumer")
		if err != nil {
			return err
		}

		loadBalancer, err := NewLoadBalancer(ctx, "pulumi-lb-as", subnets, securityGroup)
		if err != nil {
			return err
		}

		targetGroup, err := NewTargetGroup(ctx, "pulumi-tg-as", vpc)
		if err != nil {
			return err
		}

		listener, err := NewListener(ctx, "pulumi-listener-as", loadBalancer, targetGroup)
		if err != nil {
			return err
		}

		dbSubnets, err := NewDbSubnet(ctx, "dbsubnets", subnets)
		if err != nil {
			return err
		}

		bucket, err := NewBucket(ctx, "pulumi-bucket-as")
		if err != nil {
			return err
		}

		database, err := NewRDS(ctx, "pulumi-database-as", config, dbSubnets, securityGroup)
		if err != nil {
			return err
		}

		broker, err := NewRabbitMQ(ctx, "pulumi-rabbitmq-1", config, subnets)
		if err != nil {
			return err
		}

		containerDefinitionAPI := NewContainerAPI(config, database, bucket, broker, logAPI)
		taskDefinitionAPI, err := NewTask(ctx, "pulumi-task-api-as", taskExecRole, containerDefinitionAPI)
		if err != nil {
			return err
		}

		api := NewAPI(config, cluster, taskDefinitionAPI, subnets, securityGroup, targetGroup, listener)

		err = NewServiceAPI(ctx, "pulumi-api-svc-as", api)
		if err != nil {
			return err
		}

		containerDefinitionConsumer := NewContainerConsumer(config, database, bucket, broker, logConsumer)
		taskDefinitionConsumer, err := NewTask(ctx, "pulumi-task-consumer-as", taskExecRole, containerDefinitionConsumer)
		if err != nil {
			return err
		}

		consumer := NewConsumer(cluster, taskDefinitionConsumer, subnets, securityGroup)

		err = NewServiceConsumer(ctx, "pulumi-consumer-svc-as", consumer)
		if err != nil {
			return err
		}

		ctx.Export("publicHostName", loadBalancer.DnsName)
		return nil
	})
}

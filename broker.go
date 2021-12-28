package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/mq"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewRabbitMQ creates amqp broker.
func NewRabbitMQ(ctx *pulumi.Context, name string, conf Config, subnets *ec2.GetSubnetIdsResult) (*mq.Broker, error) {
	broker, err := mq.NewBroker(ctx, name, &mq.BrokerArgs{
		ApplyImmediately:   pulumi.Bool(true),
		BrokerName:         pulumi.String("pulumi-rabbitmq-1"),
		EngineType:         pulumi.String("RabbitMQ"),
		EngineVersion:      pulumi.String("3.8.23"),
		PubliclyAccessible: pulumi.Bool(true),
		HostInstanceType:   pulumi.String("mq.t3.micro"),
		SubnetIds:          getSubnet(subnets.Ids),
		Users: mq.BrokerUserArray{
			&mq.BrokerUserArgs{
				Username: conf.MqUsername,
				Password: conf.MqPassword,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return broker, nil
}

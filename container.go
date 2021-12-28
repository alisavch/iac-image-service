package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/mq"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewContainerAPI creates api container.
func NewContainerAPI(conf Config, rds *rds.Instance, bucket *s3.Bucket, broker *mq.Broker, log *cloudwatch.LogGroup) pulumi.StringOutput {
	rabbitURL := broker.Instances.Index(pulumi.Int(0)).Endpoints().Index(pulumi.Int(0))

	endpoint := split(rabbitURL)

	containerDefinition := pulumi.Sprintf(`[
		{
			"name": "%s",
			"image": "alisavch/api:latest",
			"portMappings": [
				{
					"containerPort": 8080,
					"hostPort": 8080,
					"protocol": "tcp"
				}
			],
			"environment": [
				{
					"name": "DATABASE_URL",
					"value": "postgres://%s:%s@%s/%s?sslmode=disable"
				},
				{
					"name": "BUCKET_NAME",
					"value": "%s"
				},
				{
					"name": "AWS_REGION",
					"value": "%s"
				},
				{
					"name": "AWS_ACCESS_KEY_ID",
					"value": "%s"
				},
				{
					"name": "AWS_SECRET_ACCESS_KEY",
					"value": "%s"
				},
				{
					"name": "TOKEN_TTL",
					"value": "%s"
				},
				{
					"name": "SIGNING_KEY",
					"value": "%s"
				},
				{
					"name": "REMOTE_STORAGE",
					"value": "%s"
				},
				{
					"name": "RABBITMQ_DEFAULT_USER",
					"value": "%s"
				},
				{
					"name": "RABBITMQ_DEFAULT_PASS",
					"value": "%s"
				},
				{
					"name": "RABBITMQ_URL",
					"value": "amqps://%s:%s@%s"
				}
			],
			"logConfiguration": {
                "logDriver": "awslogs",
                "options": {
                    "awslogs-group": "%s",
					"awslogs-region": "%s",
                    "awslogs-stream-prefix": "%s"
                }
			}
		}
]`, conf.APIEnv.Name,
		conf.Database.DbUsername,
		conf.Database.DbPassword,
		rds.Endpoint,
		conf.Database.DbName,
		bucket.Bucket,
		conf.Region,
		conf.AccessKeyID,
		conf.SecretAccessKey,
		conf.APIEnv.TokenLLT,
		conf.APIEnv.SigningKey,
		conf.RemoteStorage,
		conf.Broker.MqDefaultUser,
		conf.Broker.MqDefaultPassword,
		conf.Broker.MqDefaultUser,
		conf.Broker.MqDefaultPassword,
		endpoint.Index(pulumi.Int(1)),
		log.Name,
		conf.Region,
		conf.APIEnv.Name,
	)
	return containerDefinition
}

// NewContainerConsumer creates consumer container.
func NewContainerConsumer(conf Config, rds *rds.Instance, bucket *s3.Bucket, broker *mq.Broker, log *cloudwatch.LogGroup) pulumi.StringOutput {
	rabbitURL := broker.Instances.Index(pulumi.Int(0)).Endpoints().Index(pulumi.Int(0))

	endpoint := split(rabbitURL)

	containerDefinition := pulumi.Sprintf(`[
		{
			"name": "%s",
			"image": "alisavch/consumer:latest",
			"environment": [
				{
					"name": "DATABASE_URL",
					"value": "postgres://%s:%s@%s/%s?sslmode=disable"
				},
				{
					"name": "BUCKET_NAME",
					"value": "%s"
				},
				{
					"name": "AWS_REGION",
					"value": "%s"
				},
				{
					"name": "AWS_ACCESS_KEY_ID",
					"value": "%s"
				},
				{
					"name": "AWS_SECRET_ACCESS_KEY",
					"value": "%s"
				},
				{
					"name": "REMOTE_STORAGE",
					"value": "%s"
				},
				{
					"name": "RABBITMQ_DEFAULT_USER",
					"value": "%s"
				},
				{
					"name": "RABBITMQ_DEFAULT_PASS",
					"value": "%s"
				},
				{
					"name": "RABBITMQ_URL",
					"value": "amqps://%s:%s@%s"
				}
			],
			"logConfiguration": {
				"logDriver": "awslogs",
				"options": {
					"awslogs-group": "%s",
					"awslogs-region": "%s",
					"awslogs-stream-prefix": "%s"
				}
			}
		}
]`, conf.ConsumerEnv.Name,
		conf.Database.DbUsername,
		conf.Database.DbPassword,
		rds.Endpoint,
		conf.Database.DbName,
		bucket.Bucket,
		conf.Region,
		conf.AccessKeyID,
		conf.SecretAccessKey,
		conf.RemoteStorage,
		conf.Broker.MqDefaultUser,
		conf.Broker.MqDefaultPassword,
		conf.Broker.MqDefaultUser,
		conf.Broker.MqDefaultPassword,
		endpoint.Index(pulumi.Int(1)),
		log.Name,
		conf.Region,
		conf.APIEnv.Name,
	)
	return containerDefinition
}

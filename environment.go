package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Config contains the entire configuration environment.
type Config struct {
	Database
	Broker
	APIEnv
	ConsumerEnv
	Region          pulumi.String
	AccessKeyID     pulumi.StringInput
	SecretAccessKey pulumi.StringInput
	RemoteStorage   pulumi.String
}

// Database contains the database environment.
type Database struct {
	DbName     pulumi.String
	DbUsername pulumi.String
	DbPassword pulumi.StringInput
}

// Broker contains the broker environment.
type Broker struct {
	MqUsername        pulumi.String
	MqPassword        pulumi.StringInput
	MqDefaultUser     pulumi.String
	MqDefaultPassword pulumi.StringInput
}

// APIEnv contains the api environment.
type APIEnv struct {
	Name       pulumi.String
	TokenLLT   pulumi.String
	SigningKey pulumi.StringInput
}

// ConsumerEnv contains the consumer environment.
type ConsumerEnv struct {
	Name pulumi.String
}

// NewConfig configures Config.
func NewConfig(ctx *pulumi.Context) Config {
	projectConfig := config.New(ctx, "")
	awsConfig := config.New(ctx, "aws")
	region := awsConfig.Require("region")
	accessKeyID := projectConfig.RequireSecret("accessKeyID")
	secretAccessKey := projectConfig.RequireSecret("secretAccessKey")
	dbName := projectConfig.Require("dbName")
	dbUsername := projectConfig.Require("dbUsername")
	dbPassword := projectConfig.RequireSecret("dbPassword")
	mqUsername := projectConfig.Require("mqUsername")
	mqPassword := projectConfig.RequireSecret("mqPassword")
	signingKey := projectConfig.RequireSecret("signingKey")
	tokenTTL := projectConfig.Require("tokenTTL")
	remoteStorage := projectConfig.Require("remoteStorage")
	mqDefaultUser := projectConfig.Require("mqDefaultUser")
	mqDefaultPassword := projectConfig.RequireSecret("mqDefaultPassword")

	return Config{
		Database{
			DbName:     pulumi.String(dbName),
			DbUsername: pulumi.String(dbUsername),
			DbPassword: pulumi.StringInput(dbPassword),
		},
		Broker{
			MqUsername:        pulumi.String(mqUsername),
			MqPassword:        pulumi.StringInput(mqPassword),
			MqDefaultUser:     pulumi.String(mqDefaultUser),
			MqDefaultPassword: pulumi.StringInput(mqDefaultPassword),
		},
		APIEnv{
			Name:       "api",
			TokenLLT:   pulumi.String(tokenTTL),
			SigningKey: pulumi.StringInput(signingKey),
		},
		ConsumerEnv{
			Name: "consumer",
		},
		pulumi.String(region),
		pulumi.StringInput(accessKeyID),
		pulumi.StringInput(secretAccessKey),
		pulumi.String(remoteStorage),
	}
}

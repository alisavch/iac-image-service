package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewBucket creates S3 bucket.
func NewBucket(ctx *pulumi.Context, name string) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, name, &s3.BucketArgs{
		Acl: pulumi.String("private"),
		Tags: pulumi.StringMap{
			"Environment": pulumi.String("Dev"),
			"Name":        pulumi.String("pulumi-bucket-as"),
		},
	})
	if err != nil {
		return nil, err
	}
	return bucket, err
}

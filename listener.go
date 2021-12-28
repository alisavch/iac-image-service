package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewListener creates a listener that checks the connection request and the rules for routing the request.
func NewListener(ctx *pulumi.Context, name string, loadBalancer *lb.LoadBalancer, targetGroup *lb.TargetGroup) (*lb.Listener, error) {
	listener, err := lb.NewListener(ctx, name, &lb.ListenerArgs{
		LoadBalancerArn: loadBalancer.Arn,
		Port:            pulumi.Int(80),
		DefaultActions: lb.ListenerDefaultActionArray{
			lb.ListenerDefaultActionArgs{
				Type:           pulumi.String("forward"),
				TargetGroupArn: targetGroup.Arn,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return listener, nil
}

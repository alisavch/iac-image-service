package main

import (
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strings"
)

func split(input pulumi.StringOutput) pulumi.StringArrayOutput {
	return input.ApplyT(func(s string) []string {
		return strings.Split(s, "://")
	}).(pulumi.StringArrayOutput)
}

func getSubnet(a []string) pulumi.StringArrayInput {
	var res []pulumi.StringInput
	fmt.Println(a)
	res = append(res, pulumi.String(a[0]))
	return pulumi.StringArray(res)
}

// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2targetgroup


type Elasticloadbalancingv2TargetGroupMatcher struct {
	// You can specify values between 0 and 99.
	//
	// You can specify multiple values, or a range of values. The default value is 12.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_target_group#grpc_code Elasticloadbalancingv2TargetGroup#grpc_code}
	GrpcCode *string `field:"optional" json:"grpcCode" yaml:"grpcCode"`
	// For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200.
	//
	// You can specify multiple values or a range of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_target_group#http_code Elasticloadbalancingv2TargetGroup#http_code}
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
}


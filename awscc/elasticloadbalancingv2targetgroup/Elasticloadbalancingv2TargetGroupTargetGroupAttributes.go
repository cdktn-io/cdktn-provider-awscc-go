// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2targetgroup


type Elasticloadbalancingv2TargetGroupTargetGroupAttributes struct {
	// The value of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#key Elasticloadbalancingv2TargetGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The name of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_target_group#value Elasticloadbalancingv2TargetGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


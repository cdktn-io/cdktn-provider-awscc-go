// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleTags struct {
	// The key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_listener_rule#key Elasticloadbalancingv2ListenerRule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_listener_rule#value Elasticloadbalancingv2ListenerRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


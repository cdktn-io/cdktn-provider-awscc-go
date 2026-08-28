// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleActionsForwardConfigTargetGroups struct {
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener_rule#target_group_arn Elasticloadbalancingv2ListenerRule#target_group_arn}
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
	// The weight. The range is 0 to 999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_listener_rule#weight Elasticloadbalancingv2ListenerRule#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleActionsForwardConfig struct {
	// Information about how traffic will be distributed between multiple target groups in a forward rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener_rule#target_groups Elasticloadbalancingv2ListenerRule#target_groups}
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// Information about the target group stickiness for a rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener_rule#target_group_stickiness_config Elasticloadbalancingv2ListenerRule#target_group_stickiness_config}
	TargetGroupStickinessConfig *Elasticloadbalancingv2ListenerRuleActionsForwardConfigTargetGroupStickinessConfig `field:"optional" json:"targetGroupStickinessConfig" yaml:"targetGroupStickinessConfig"`
}


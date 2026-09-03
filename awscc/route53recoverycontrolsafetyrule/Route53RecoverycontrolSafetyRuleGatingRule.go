// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoverycontrolsafetyrule


type Route53RecoverycontrolSafetyRuleGatingRule struct {
	// The gating controls for the gating rule.
	//
	// That is, routing controls that are evaluated by the rule configuration that you specify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoverycontrol_safety_rule#gating_controls Route53RecoverycontrolSafetyRule#gating_controls}
	GatingControls *[]*string `field:"optional" json:"gatingControls" yaml:"gatingControls"`
	// Routing controls that can only be set or unset if the specified RuleConfig evaluates to true for the specified GatingControls.
	//
	// For example, say you have three gating controls, one for each of three AWS Regions. Now you specify AtLeast 2 as your RuleConfig. With these settings, you can only change (set or unset) the routing controls that you have specified as TargetControls if that rule evaluates to true.
	// In other words, your ability to change the routing controls that you have specified as TargetControls is gated by the rule that you set for the routing controls in GatingControls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoverycontrol_safety_rule#target_controls Route53RecoverycontrolSafetyRule#target_controls}
	TargetControls *[]*string `field:"optional" json:"targetControls" yaml:"targetControls"`
	// An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail.
	//
	// This helps prevent "flapping" of state. The wait period is 5000 ms by default, but you can choose a custom value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoverycontrol_safety_rule#wait_period_ms Route53RecoverycontrolSafetyRule#wait_period_ms}
	WaitPeriodMs *float64 `field:"optional" json:"waitPeriodMs" yaml:"waitPeriodMs"`
}


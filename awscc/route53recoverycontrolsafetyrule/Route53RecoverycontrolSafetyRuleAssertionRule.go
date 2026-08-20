// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoverycontrolsafetyrule


type Route53RecoverycontrolSafetyRuleAssertionRule struct {
	// The routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed.
	//
	// For example, you might include three routing controls, one for each of three AWS Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53recoverycontrol_safety_rule#asserted_controls Route53RecoverycontrolSafetyRule#asserted_controls}
	AssertedControls *[]*string `field:"optional" json:"assertedControls" yaml:"assertedControls"`
	// An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail.
	//
	// This helps prevent "flapping" of state. The wait period is 5000 ms by default, but you can choose a custom value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53recoverycontrol_safety_rule#wait_period_ms Route53RecoverycontrolSafetyRule#wait_period_ms}
	WaitPeriodMs *float64 `field:"optional" json:"waitPeriodMs" yaml:"waitPeriodMs"`
}


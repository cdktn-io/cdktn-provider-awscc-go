// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoverycontrolsafetyrule


type Route53RecoverycontrolSafetyRuleRuleConfig struct {
	// Logical negation of the rule. If the rule would usually evaluate true, it's evaluated as false, and vice versa.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoverycontrol_safety_rule#inverted Route53RecoverycontrolSafetyRule#inverted}
	Inverted interface{} `field:"optional" json:"inverted" yaml:"inverted"`
	// The value of N, when you specify an ATLEAST rule type.
	//
	// That is, Threshold is the number of controls that must be set when you specify an ATLEAST type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoverycontrol_safety_rule#threshold Route53RecoverycontrolSafetyRule#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// A rule can be one of the following: ATLEAST, AND, or OR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53recoverycontrol_safety_rule#type Route53RecoverycontrolSafetyRule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


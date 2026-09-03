// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewrites struct {
	// The regular expression to match in the input string. The maximum length of the string is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticloadbalancingv2_listener_rule#regex Elasticloadbalancingv2ListenerRule#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// The replacement string to use when rewriting the matched input.
	//
	// The maximum length of the string is 1,024 characters. You can specify capture groups in the regular expression (for example, $1 and $2).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticloadbalancingv2_listener_rule#replace Elasticloadbalancingv2ListenerRule#replace}
	Replace *string `field:"optional" json:"replace" yaml:"replace"`
}


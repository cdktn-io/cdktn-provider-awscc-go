// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigrule


type ConfigConfigRuleSourceCustomPolicyDetails struct {
	// The boolean expression for enabling debug logging for your CC Custom Policy rule. The default value is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_config_rule#enable_debug_log_delivery ConfigConfigRule#enable_debug_log_delivery}
	EnableDebugLogDelivery interface{} `field:"optional" json:"enableDebugLogDelivery" yaml:"enableDebugLogDelivery"`
	// The runtime system for your CC Custom Policy rule.
	//
	// Guard is a policy-as-code language that allows you to write policies that are enforced by CC Custom Policy rules. For more information about Guard, see the [Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_config_rule#policy_runtime ConfigConfigRule#policy_runtime}
	PolicyRuntime *string `field:"optional" json:"policyRuntime" yaml:"policyRuntime"`
	// The policy definition containing the logic for your CC Custom Policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_config_rule#policy_text ConfigConfigRule#policy_text}
	PolicyText *string `field:"optional" json:"policyText" yaml:"policyText"`
}


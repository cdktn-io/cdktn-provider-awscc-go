// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigrule


type ConfigConfigRuleEvaluationModes struct {
	// The mode of an evaluation. The valid values are Detective or Proactive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/config_config_rule#mode ConfigConfigRule#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}


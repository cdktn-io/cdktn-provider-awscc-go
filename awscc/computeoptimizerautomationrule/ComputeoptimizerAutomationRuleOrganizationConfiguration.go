// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerautomationrule


type ComputeoptimizerAutomationRuleOrganizationConfiguration struct {
	// List of account IDs where the organization rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#account_ids ComputeoptimizerAutomationRule#account_ids}
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// When the rule should be applied relative to account rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#rule_apply_order ComputeoptimizerAutomationRule#rule_apply_order}
	RuleApplyOrder *string `field:"optional" json:"ruleApplyOrder" yaml:"ruleApplyOrder"`
}


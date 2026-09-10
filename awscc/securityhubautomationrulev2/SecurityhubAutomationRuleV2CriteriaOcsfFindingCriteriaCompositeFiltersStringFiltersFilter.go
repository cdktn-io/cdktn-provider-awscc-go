// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersStringFiltersFilter struct {
	// The condition to apply to a string value when filtering findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#comparison SecurityhubAutomationRuleV2#comparison}
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// The string filter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#value SecurityhubAutomationRuleV2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


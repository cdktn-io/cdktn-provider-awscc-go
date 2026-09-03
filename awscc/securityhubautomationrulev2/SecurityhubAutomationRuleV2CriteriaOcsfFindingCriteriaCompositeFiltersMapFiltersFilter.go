// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersMapFiltersFilter struct {
	// The condition to apply to the key value when filtering findings with a map filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_automation_rule_v2#comparison SecurityhubAutomationRuleV2#comparison}
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// The key of the map filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_automation_rule_v2#key SecurityhubAutomationRuleV2#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the key in the map filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_automation_rule_v2#value SecurityhubAutomationRuleV2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersDateFiltersFilterDateRange struct {
	// A date range unit for the date filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_automation_rule_v2#unit SecurityhubAutomationRuleV2#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// A date range value for the date filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_automation_rule_v2#value SecurityhubAutomationRuleV2#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersDateFiltersFilter struct {
	// A date range for the date filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_automation_rule_v2#date_range SecurityhubAutomationRuleV2#date_range}
	DateRange *SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersDateFiltersFilterDateRange `field:"optional" json:"dateRange" yaml:"dateRange"`
	// The timestamp formatted in ISO8601.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_automation_rule_v2#end SecurityhubAutomationRuleV2#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// The timestamp formatted in ISO8601.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_automation_rule_v2#start SecurityhubAutomationRuleV2#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
}


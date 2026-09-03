// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersNumberFiltersFilter struct {
	// The equal-to condition to be applied to a single field when querying for findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_automation_rule_v2#eq SecurityhubAutomationRuleV2#eq}
	Eq *float64 `field:"optional" json:"eq" yaml:"eq"`
	// The greater-than-equal condition to be applied to a single field when querying for findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_automation_rule_v2#gte SecurityhubAutomationRuleV2#gte}
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// The less-than-equal condition to be applied to a single field when querying for findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_automation_rule_v2#lte SecurityhubAutomationRuleV2#lte}
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
}


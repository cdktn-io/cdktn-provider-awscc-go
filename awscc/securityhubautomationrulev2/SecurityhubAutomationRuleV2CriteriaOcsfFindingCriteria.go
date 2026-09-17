// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteria struct {
	// Enables the creation of complex filtering conditions by combining filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_automation_rule_v2#composite_filters SecurityhubAutomationRuleV2#composite_filters}
	CompositeFilters interface{} `field:"optional" json:"compositeFilters" yaml:"compositeFilters"`
	// The logical operator used to combine multiple conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_automation_rule_v2#composite_operator SecurityhubAutomationRuleV2#composite_operator}
	CompositeOperator *string `field:"optional" json:"compositeOperator" yaml:"compositeOperator"`
}


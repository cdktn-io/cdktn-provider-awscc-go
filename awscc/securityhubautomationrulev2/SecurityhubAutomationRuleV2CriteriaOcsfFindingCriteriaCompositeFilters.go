// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFilters struct {
	// Enables filtering based on boolean field values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_automation_rule_v2#boolean_filters SecurityhubAutomationRuleV2#boolean_filters}
	BooleanFilters interface{} `field:"optional" json:"booleanFilters" yaml:"booleanFilters"`
	// Enables filtering based on date and timestamp fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_automation_rule_v2#date_filters SecurityhubAutomationRuleV2#date_filters}
	DateFilters interface{} `field:"optional" json:"dateFilters" yaml:"dateFilters"`
	// Enables filtering based on map field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_automation_rule_v2#map_filters SecurityhubAutomationRuleV2#map_filters}
	MapFilters interface{} `field:"optional" json:"mapFilters" yaml:"mapFilters"`
	// Enables filtering based on numerical field values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_automation_rule_v2#number_filters SecurityhubAutomationRuleV2#number_filters}
	NumberFilters interface{} `field:"optional" json:"numberFilters" yaml:"numberFilters"`
	// The logical operator used to combine multiple conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_automation_rule_v2#operator SecurityhubAutomationRuleV2#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Enables filtering based on string field values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_automation_rule_v2#string_filters SecurityhubAutomationRuleV2#string_filters}
	StringFilters interface{} `field:"optional" json:"stringFilters" yaml:"stringFilters"`
}


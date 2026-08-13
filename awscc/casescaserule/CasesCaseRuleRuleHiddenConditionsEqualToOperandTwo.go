// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescaserule


type CasesCaseRuleRuleHiddenConditionsEqualToOperandTwo struct {
	// A boolean value to compare against the field value in the condition evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cases_case_rule#boolean_value CasesCaseRule#boolean_value}
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// A numeric value to compare against the field value in the condition evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cases_case_rule#double_value CasesCaseRule#double_value}
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// An empty operand value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cases_case_rule#empty_value CasesCaseRule#empty_value}
	EmptyValue *string `field:"optional" json:"emptyValue" yaml:"emptyValue"`
	// A string value to compare against the field value in the condition evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cases_case_rule#string_value CasesCaseRule#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}


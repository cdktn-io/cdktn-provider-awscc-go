// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescaserule


type CasesCaseRuleRuleHiddenConditions struct {
	// Boolean operands for a condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cases_case_rule#equal_to CasesCaseRule#equal_to}
	EqualTo *CasesCaseRuleRuleHiddenConditionsEqualTo `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Boolean operands for a condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cases_case_rule#not_equal_to CasesCaseRule#not_equal_to}
	NotEqualTo *CasesCaseRuleRuleHiddenConditionsNotEqualTo `field:"optional" json:"notEqualTo" yaml:"notEqualTo"`
}


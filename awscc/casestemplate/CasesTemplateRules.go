// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casestemplate


type CasesTemplateRules struct {
	// The unique identifier of a case rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_template#case_rule_id CasesTemplate#case_rule_id}
	CaseRuleId *string `field:"optional" json:"caseRuleId" yaml:"caseRuleId"`
	// The ID of the field that this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_template#field_id CasesTemplate#field_id}
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}


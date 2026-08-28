// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescaserule


type CasesCaseRuleRuleRequiredConditionsNotEqualTo struct {
	// The left hand operand in the condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cases_case_rule#operand_one CasesCaseRule#operand_one}
	OperandOne *CasesCaseRuleRuleRequiredConditionsNotEqualToOperandOne `field:"optional" json:"operandOne" yaml:"operandOne"`
	// The right hand operand in the condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cases_case_rule#operand_two CasesCaseRule#operand_two}
	OperandTwo *CasesCaseRuleRuleRequiredConditionsNotEqualToOperandTwo `field:"optional" json:"operandTwo" yaml:"operandTwo"`
	// The value of the outer rule if the condition evaluates to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cases_case_rule#result CasesCaseRule#result}
	Result interface{} `field:"optional" json:"result" yaml:"result"`
}


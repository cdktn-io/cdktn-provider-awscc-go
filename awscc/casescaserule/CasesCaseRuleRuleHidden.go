// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescaserule


type CasesCaseRuleRuleHidden struct {
	// List of conditions for the hidden rule;
	//
	// the first condition to evaluate to true dictates the value of the rule
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cases_case_rule#conditions CasesCaseRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The value of the rule (i.e. whether the field is hidden) should none of the conditions evaluate to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cases_case_rule#default_value CasesCaseRule#default_value}
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}


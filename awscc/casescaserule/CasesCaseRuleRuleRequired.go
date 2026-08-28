// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescaserule


type CasesCaseRuleRuleRequired struct {
	// An ordered list of boolean conditions that determine when the field should be required.
	//
	// Conditions are evaluated in order, and the first condition that evaluates to true determines whether the field is required, overriding the default value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cases_case_rule#conditions CasesCaseRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The default required state for the field when none of the specified conditions are met.
	//
	// If true, the field is required by default; if false, the field is optional by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cases_case_rule#default_value CasesCaseRule#default_value}
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}


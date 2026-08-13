// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescaserule


type CasesCaseRuleRule struct {
	// Hidden rule type, used to indicate whether a field is hidden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cases_case_rule#hidden CasesCaseRule#hidden}
	Hidden *CasesCaseRuleRuleHidden `field:"optional" json:"hidden" yaml:"hidden"`
	// A required rule type, used to indicate whether a field is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cases_case_rule#required CasesCaseRule#required}
	Required *CasesCaseRuleRuleRequired `field:"optional" json:"required" yaml:"required"`
}


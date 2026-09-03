// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesConditionsStringExpressionEvaluateAnalysis struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ses_mail_manager_rule_set#analyzer SesMailManagerRuleSet#analyzer}.
	Analyzer *string `field:"optional" json:"analyzer" yaml:"analyzer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ses_mail_manager_rule_set#result_field SesMailManagerRuleSet#result_field}.
	ResultField *string `field:"optional" json:"resultField" yaml:"resultField"`
}


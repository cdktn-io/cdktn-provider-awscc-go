// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesUnlessVerdictExpressionEvaluate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_rule_set#analysis SesMailManagerRuleSet#analysis}.
	Analysis *SesMailManagerRuleSetRulesUnlessVerdictExpressionEvaluateAnalysis `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_rule_set#attribute SesMailManagerRuleSet#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}


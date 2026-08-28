// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesConditionsStringExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#evaluate SesMailManagerRuleSet#evaluate}.
	Evaluate *SesMailManagerRuleSetRulesConditionsStringExpressionEvaluate `field:"optional" json:"evaluate" yaml:"evaluate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#operator SesMailManagerRuleSet#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#values SesMailManagerRuleSet#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


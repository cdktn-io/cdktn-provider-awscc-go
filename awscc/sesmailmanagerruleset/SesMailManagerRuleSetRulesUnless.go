// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesUnless struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_rule_set#boolean_expression SesMailManagerRuleSet#boolean_expression}.
	BooleanExpression *SesMailManagerRuleSetRulesUnlessBooleanExpression `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_rule_set#dmarc_expression SesMailManagerRuleSet#dmarc_expression}.
	DmarcExpression *SesMailManagerRuleSetRulesUnlessDmarcExpression `field:"optional" json:"dmarcExpression" yaml:"dmarcExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_rule_set#ip_expression SesMailManagerRuleSet#ip_expression}.
	IpExpression *SesMailManagerRuleSetRulesUnlessIpExpression `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_rule_set#number_expression SesMailManagerRuleSet#number_expression}.
	NumberExpression *SesMailManagerRuleSetRulesUnlessNumberExpression `field:"optional" json:"numberExpression" yaml:"numberExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_rule_set#string_expression SesMailManagerRuleSet#string_expression}.
	StringExpression *SesMailManagerRuleSetRulesUnlessStringExpression `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_rule_set#verdict_expression SesMailManagerRuleSet#verdict_expression}.
	VerdictExpression *SesMailManagerRuleSetRulesUnlessVerdictExpression `field:"optional" json:"verdictExpression" yaml:"verdictExpression"`
}


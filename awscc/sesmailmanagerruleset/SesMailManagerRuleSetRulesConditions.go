// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_mail_manager_rule_set#boolean_expression SesMailManagerRuleSet#boolean_expression}.
	BooleanExpression *SesMailManagerRuleSetRulesConditionsBooleanExpression `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_mail_manager_rule_set#dmarc_expression SesMailManagerRuleSet#dmarc_expression}.
	DmarcExpression *SesMailManagerRuleSetRulesConditionsDmarcExpression `field:"optional" json:"dmarcExpression" yaml:"dmarcExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_mail_manager_rule_set#ip_expression SesMailManagerRuleSet#ip_expression}.
	IpExpression *SesMailManagerRuleSetRulesConditionsIpExpression `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_mail_manager_rule_set#number_expression SesMailManagerRuleSet#number_expression}.
	NumberExpression *SesMailManagerRuleSetRulesConditionsNumberExpression `field:"optional" json:"numberExpression" yaml:"numberExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_mail_manager_rule_set#string_expression SesMailManagerRuleSet#string_expression}.
	StringExpression *SesMailManagerRuleSetRulesConditionsStringExpression `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_mail_manager_rule_set#verdict_expression SesMailManagerRuleSet#verdict_expression}.
	VerdictExpression *SesMailManagerRuleSetRulesConditionsVerdictExpression `field:"optional" json:"verdictExpression" yaml:"verdictExpression"`
}


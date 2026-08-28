// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesActionsRelay struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#action_failure_policy SesMailManagerRuleSet#action_failure_policy}.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#mail_from SesMailManagerRuleSet#mail_from}.
	MailFrom *string `field:"optional" json:"mailFrom" yaml:"mailFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#relay SesMailManagerRuleSet#relay}.
	Relay *string `field:"optional" json:"relay" yaml:"relay"`
}


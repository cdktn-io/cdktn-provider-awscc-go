// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesActionsBounce struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#action_failure_policy SesMailManagerRuleSet#action_failure_policy}.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#diagnostic_message SesMailManagerRuleSet#diagnostic_message}.
	DiagnosticMessage *string `field:"optional" json:"diagnosticMessage" yaml:"diagnosticMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#message SesMailManagerRuleSet#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#role_arn SesMailManagerRuleSet#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#sender SesMailManagerRuleSet#sender}.
	Sender *string `field:"optional" json:"sender" yaml:"sender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#smtp_reply_code SesMailManagerRuleSet#smtp_reply_code}.
	SmtpReplyCode *string `field:"optional" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_mail_manager_rule_set#status_code SesMailManagerRuleSet#status_code}.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}


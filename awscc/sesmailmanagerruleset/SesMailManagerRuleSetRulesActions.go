// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerruleset


type SesMailManagerRuleSetRulesActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#add_header SesMailManagerRuleSet#add_header}.
	AddHeader *SesMailManagerRuleSetRulesActionsAddHeader `field:"optional" json:"addHeader" yaml:"addHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#archive SesMailManagerRuleSet#archive}.
	Archive *SesMailManagerRuleSetRulesActionsArchive `field:"optional" json:"archive" yaml:"archive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#bounce SesMailManagerRuleSet#bounce}.
	Bounce *SesMailManagerRuleSetRulesActionsBounce `field:"optional" json:"bounce" yaml:"bounce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#deliver_to_mailbox SesMailManagerRuleSet#deliver_to_mailbox}.
	DeliverToMailbox *SesMailManagerRuleSetRulesActionsDeliverToMailbox `field:"optional" json:"deliverToMailbox" yaml:"deliverToMailbox"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#deliver_to_q_business SesMailManagerRuleSet#deliver_to_q_business}.
	DeliverToQBusiness *SesMailManagerRuleSetRulesActionsDeliverToQBusiness `field:"optional" json:"deliverToQBusiness" yaml:"deliverToQBusiness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#drop SesMailManagerRuleSet#drop}.
	Drop *string `field:"optional" json:"drop" yaml:"drop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#invoke_lambda SesMailManagerRuleSet#invoke_lambda}.
	InvokeLambda *SesMailManagerRuleSetRulesActionsInvokeLambda `field:"optional" json:"invokeLambda" yaml:"invokeLambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#publish_to_sns SesMailManagerRuleSet#publish_to_sns}.
	PublishToSns *SesMailManagerRuleSetRulesActionsPublishToSns `field:"optional" json:"publishToSns" yaml:"publishToSns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#relay SesMailManagerRuleSet#relay}.
	Relay *SesMailManagerRuleSetRulesActionsRelay `field:"optional" json:"relay" yaml:"relay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#replace_recipient SesMailManagerRuleSet#replace_recipient}.
	ReplaceRecipient *SesMailManagerRuleSetRulesActionsReplaceRecipient `field:"optional" json:"replaceRecipient" yaml:"replaceRecipient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#send SesMailManagerRuleSet#send}.
	Send *SesMailManagerRuleSetRulesActionsSend `field:"optional" json:"send" yaml:"send"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_rule_set#write_to_s3 SesMailManagerRuleSet#write_to_s3}.
	WriteToS3 *SesMailManagerRuleSetRulesActionsWriteToS3 `field:"optional" json:"writeToS3" yaml:"writeToS3"`
}


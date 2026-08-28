// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolriskconfigurationattachment


type CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationNotifyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#block_email CognitoUserPoolRiskConfigurationAttachment#block_email}.
	BlockEmail *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationNotifyConfigurationBlockEmail `field:"optional" json:"blockEmail" yaml:"blockEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#from CognitoUserPoolRiskConfigurationAttachment#from}.
	From *string `field:"optional" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#mfa_email CognitoUserPoolRiskConfigurationAttachment#mfa_email}.
	MfaEmail *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationNotifyConfigurationMfaEmail `field:"optional" json:"mfaEmail" yaml:"mfaEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#no_action_email CognitoUserPoolRiskConfigurationAttachment#no_action_email}.
	NoActionEmail *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationNotifyConfigurationNoActionEmail `field:"optional" json:"noActionEmail" yaml:"noActionEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#reply_to CognitoUserPoolRiskConfigurationAttachment#reply_to}.
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#source_arn CognitoUserPoolRiskConfigurationAttachment#source_arn}.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolriskconfigurationattachment


type CognitoUserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#actions CognitoUserPoolRiskConfigurationAttachment#actions}.
	Actions *CognitoUserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationActions `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#event_filter CognitoUserPoolRiskConfigurationAttachment#event_filter}.
	EventFilter *[]*string `field:"optional" json:"eventFilter" yaml:"eventFilter"`
}


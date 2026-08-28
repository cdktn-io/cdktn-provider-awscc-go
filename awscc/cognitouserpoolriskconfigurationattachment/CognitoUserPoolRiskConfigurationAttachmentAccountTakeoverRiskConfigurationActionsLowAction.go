// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolriskconfigurationattachment


type CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationActionsLowAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#event_action CognitoUserPoolRiskConfigurationAttachment#event_action}.
	EventAction *string `field:"optional" json:"eventAction" yaml:"eventAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_risk_configuration_attachment#notify CognitoUserPoolRiskConfigurationAttachment#notify}.
	Notify interface{} `field:"optional" json:"notify" yaml:"notify"`
}


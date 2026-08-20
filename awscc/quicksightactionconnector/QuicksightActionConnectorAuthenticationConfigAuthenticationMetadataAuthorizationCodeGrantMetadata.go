// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightactionconnector


type QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataAuthorizationCodeGrantMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_action_connector#authorization_code_grant_credentials_details QuicksightActionConnector#authorization_code_grant_credentials_details}.
	AuthorizationCodeGrantCredentialsDetails *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataAuthorizationCodeGrantMetadataAuthorizationCodeGrantCredentialsDetails `field:"optional" json:"authorizationCodeGrantCredentialsDetails" yaml:"authorizationCodeGrantCredentialsDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_action_connector#authorization_code_grant_credentials_source QuicksightActionConnector#authorization_code_grant_credentials_source}.
	AuthorizationCodeGrantCredentialsSource *string `field:"optional" json:"authorizationCodeGrantCredentialsSource" yaml:"authorizationCodeGrantCredentialsSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_action_connector#base_endpoint QuicksightActionConnector#base_endpoint}.
	BaseEndpoint *string `field:"optional" json:"baseEndpoint" yaml:"baseEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_action_connector#redirect_url QuicksightActionConnector#redirect_url}.
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
}


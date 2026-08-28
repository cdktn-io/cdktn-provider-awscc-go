// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightactionconnector


type QuicksightActionConnectorAuthenticationConfigAuthenticationMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_action_connector#api_key_connection_metadata QuicksightActionConnector#api_key_connection_metadata}.
	ApiKeyConnectionMetadata *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataApiKeyConnectionMetadata `field:"optional" json:"apiKeyConnectionMetadata" yaml:"apiKeyConnectionMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_action_connector#authorization_code_grant_metadata QuicksightActionConnector#authorization_code_grant_metadata}.
	AuthorizationCodeGrantMetadata *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataAuthorizationCodeGrantMetadata `field:"optional" json:"authorizationCodeGrantMetadata" yaml:"authorizationCodeGrantMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_action_connector#basic_auth_connection_metadata QuicksightActionConnector#basic_auth_connection_metadata}.
	BasicAuthConnectionMetadata *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataBasicAuthConnectionMetadata `field:"optional" json:"basicAuthConnectionMetadata" yaml:"basicAuthConnectionMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_action_connector#client_credentials_grant_metadata QuicksightActionConnector#client_credentials_grant_metadata}.
	ClientCredentialsGrantMetadata *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataClientCredentialsGrantMetadata `field:"optional" json:"clientCredentialsGrantMetadata" yaml:"clientCredentialsGrantMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_action_connector#iam_connection_metadata QuicksightActionConnector#iam_connection_metadata}.
	IamConnectionMetadata *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataIamConnectionMetadata `field:"optional" json:"iamConnectionMetadata" yaml:"iamConnectionMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_action_connector#none_connection_metadata QuicksightActionConnector#none_connection_metadata}.
	NoneConnectionMetadata *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataNoneConnectionMetadata `field:"optional" json:"noneConnectionMetadata" yaml:"noneConnectionMetadata"`
}


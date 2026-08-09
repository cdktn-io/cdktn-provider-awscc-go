// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightactionconnector


type QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataClientCredentialsGrantMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_action_connector#base_endpoint QuicksightActionConnector#base_endpoint}.
	BaseEndpoint *string `field:"optional" json:"baseEndpoint" yaml:"baseEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_action_connector#client_credentials_details QuicksightActionConnector#client_credentials_details}.
	ClientCredentialsDetails *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataClientCredentialsGrantMetadataClientCredentialsDetails `field:"optional" json:"clientCredentialsDetails" yaml:"clientCredentialsDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_action_connector#client_credentials_source QuicksightActionConnector#client_credentials_source}.
	ClientCredentialsSource *string `field:"optional" json:"clientCredentialsSource" yaml:"clientCredentialsSource"`
}


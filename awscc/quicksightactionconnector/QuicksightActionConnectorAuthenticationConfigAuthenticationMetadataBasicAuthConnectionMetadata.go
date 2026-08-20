// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightactionconnector


type QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataBasicAuthConnectionMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_action_connector#base_endpoint QuicksightActionConnector#base_endpoint}.
	BaseEndpoint *string `field:"optional" json:"baseEndpoint" yaml:"baseEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_action_connector#password QuicksightActionConnector#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_action_connector#username QuicksightActionConnector#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}


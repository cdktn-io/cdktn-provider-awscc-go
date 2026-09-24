// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOAuthParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#identity_provider_resource_uri QuicksightDataSource#identity_provider_resource_uri}.
	IdentityProviderResourceUri *string `field:"optional" json:"identityProviderResourceUri" yaml:"identityProviderResourceUri"`
	// <p>VPC connection properties.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#identity_provider_vpc_connection_properties QuicksightDataSource#identity_provider_vpc_connection_properties}
	IdentityProviderVpcConnectionProperties *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOAuthParametersIdentityProviderVpcConnectionProperties `field:"optional" json:"identityProviderVpcConnectionProperties" yaml:"identityProviderVpcConnectionProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#o_auth_scope QuicksightDataSource#o_auth_scope}.
	OAuthScope *string `field:"optional" json:"oAuthScope" yaml:"oAuthScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#token_provider_url QuicksightDataSource#token_provider_url}.
	TokenProviderUrl *string `field:"optional" json:"tokenProviderUrl" yaml:"tokenProviderUrl"`
}


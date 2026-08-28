// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationTokenSelection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_identity_source#access_token_only VerifiedpermissionsIdentitySource#access_token_only}.
	AccessTokenOnly *VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationTokenSelectionAccessTokenOnly `field:"optional" json:"accessTokenOnly" yaml:"accessTokenOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_identity_source#identity_token_only VerifiedpermissionsIdentitySource#identity_token_only}.
	IdentityTokenOnly *VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationTokenSelectionIdentityTokenOnly `field:"optional" json:"identityTokenOnly" yaml:"identityTokenOnly"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/verifiedpermissions_identity_source#entity_id_prefix VerifiedpermissionsIdentitySource#entity_id_prefix}.
	EntityIdPrefix *string `field:"optional" json:"entityIdPrefix" yaml:"entityIdPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/verifiedpermissions_identity_source#group_configuration VerifiedpermissionsIdentitySource#group_configuration}.
	GroupConfiguration *VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationGroupConfiguration `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/verifiedpermissions_identity_source#issuer VerifiedpermissionsIdentitySource#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/verifiedpermissions_identity_source#token_selection VerifiedpermissionsIdentitySource#token_selection}.
	TokenSelection *VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationTokenSelection `field:"optional" json:"tokenSelection" yaml:"tokenSelection"`
}


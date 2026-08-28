// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfigurationCognitoUserPoolConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_identity_source#client_ids VerifiedpermissionsIdentitySource#client_ids}.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_identity_source#group_configuration VerifiedpermissionsIdentitySource#group_configuration}.
	GroupConfiguration *VerifiedpermissionsIdentitySourceConfigurationCognitoUserPoolConfigurationGroupConfiguration `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/verifiedpermissions_identity_source#user_pool_arn VerifiedpermissionsIdentitySource#user_pool_arn}.
	UserPoolArn *string `field:"optional" json:"userPoolArn" yaml:"userPoolArn"`
}


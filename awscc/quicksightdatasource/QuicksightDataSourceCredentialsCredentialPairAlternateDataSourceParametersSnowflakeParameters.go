// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#authentication_type QuicksightDataSource#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// <p>Database.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#database_access_control_role QuicksightDataSource#database_access_control_role}.
	DatabaseAccessControlRole *string `field:"optional" json:"databaseAccessControlRole" yaml:"databaseAccessControlRole"`
	// <p>Host.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#o_auth_parameters QuicksightDataSource#o_auth_parameters}.
	OAuthParameters *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersSnowflakeParametersOAuthParameters `field:"optional" json:"oAuthParameters" yaml:"oAuthParameters"`
	// <p>Warehouse.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_source#warehouse QuicksightDataSource#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}


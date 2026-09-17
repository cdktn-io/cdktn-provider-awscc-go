// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceDataSourceParametersStarburstParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#authentication_type QuicksightDataSource#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// <p>The catalog name for the Starburst data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#catalog QuicksightDataSource#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#database_access_control_role QuicksightDataSource#database_access_control_role}.
	DatabaseAccessControlRole *string `field:"optional" json:"databaseAccessControlRole" yaml:"databaseAccessControlRole"`
	// <p>The host name of the Starburst data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#o_auth_parameters QuicksightDataSource#o_auth_parameters}.
	OAuthParameters *QuicksightDataSourceDataSourceParametersStarburstParametersOAuthParameters `field:"optional" json:"oAuthParameters" yaml:"oAuthParameters"`
	// <p>The port for the Starburst data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#product_type QuicksightDataSource#product_type}.
	ProductType *string `field:"optional" json:"productType" yaml:"productType"`
}


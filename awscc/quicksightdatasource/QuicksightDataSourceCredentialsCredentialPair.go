// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceCredentialsCredentialPair struct {
	// <p>A set of alternate data source parameters that you want to share for these             credentials.
	//
	// The credentials are applied in tandem with the data source parameters when
	//             you copy a data source by using a create or update request. The API operation compares
	//             the <code>DataSourceParameters</code> structure that's in the request with the
	//             structures in the <code>AlternateDataSourceParameters</code> allow list. If the
	//             structures are an exact match, the request is allowed to use the new data source with
	//             the existing credentials. If the <code>AlternateDataSourceParameters</code> list is
	//             null, the <code>DataSourceParameters</code> originally used with these
	//                 <code>Credentials</code> is automatically allowed.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_source#alternate_data_source_parameters QuicksightDataSource#alternate_data_source_parameters}
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
	// <p>Password.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_source#password QuicksightDataSource#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// <p>User name.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_source#username QuicksightDataSource#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}


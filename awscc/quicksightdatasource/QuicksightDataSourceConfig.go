// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#name QuicksightDataSource#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#type QuicksightDataSource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// <p>A set of alternate data source parameters that you want to share for the credentials             stored with this data source.
	//
	// The credentials are applied in tandem with the data source
	//             parameters when you copy a data source by using a create or update request. The API
	//             operation compares the <code>DataSourceParameters</code> structure that's in the request
	//             with the structures in the <code>AlternateDataSourceParameters</code> allow list. If the
	//             structures are an exact match, the request is allowed to use the credentials from this
	//             existing data source. If the <code>AlternateDataSourceParameters</code> list is null,
	//             the <code>Credentials</code> originally used with this <code>DataSourceParameters</code>
	//             are automatically allowed.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#alternate_data_source_parameters QuicksightDataSource#alternate_data_source_parameters}
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#aws_account_id QuicksightDataSource#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// <p>Data source credentials.
	//
	// This is a variant type structure. For this structure to be
	//             valid, only one of the attributes can be non-null.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#credentials QuicksightDataSource#credentials}
	Credentials *QuicksightDataSourceCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#data_source_id QuicksightDataSource#data_source_id}.
	DataSourceId *string `field:"optional" json:"dataSourceId" yaml:"dataSourceId"`
	// <p>The parameters that Amazon QuickSight uses to connect to your underlying data source.
	//
	// This is a variant type structure. For this structure to be valid, only one of the
	//             attributes can be non-null.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#data_source_parameters QuicksightDataSource#data_source_parameters}
	DataSourceParameters *QuicksightDataSourceDataSourceParameters `field:"optional" json:"dataSourceParameters" yaml:"dataSourceParameters"`
	// <p>Error information for the data source creation or update.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#error_info QuicksightDataSource#error_info}
	ErrorInfo *QuicksightDataSourceErrorInfo `field:"optional" json:"errorInfo" yaml:"errorInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#folder_arns QuicksightDataSource#folder_arns}.
	FolderArns *[]*string `field:"optional" json:"folderArns" yaml:"folderArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#permissions QuicksightDataSource#permissions}.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// <p>Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your             underlying data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#ssl_properties QuicksightDataSource#ssl_properties}
	SslProperties *QuicksightDataSourceSslProperties `field:"optional" json:"sslProperties" yaml:"sslProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#tags QuicksightDataSource#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// <p>VPC connection properties.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_source#vpc_connection_properties QuicksightDataSource#vpc_connection_properties}
	VpcConnectionProperties *QuicksightDataSourceVpcConnectionProperties `field:"optional" json:"vpcConnectionProperties" yaml:"vpcConnectionProperties"`
}


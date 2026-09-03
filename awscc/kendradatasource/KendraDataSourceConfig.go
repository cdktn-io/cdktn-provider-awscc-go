// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraDataSourceConfig struct {
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
	// ID of Index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#index_id KendraDataSource#index_id}
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// Name of data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#name KendraDataSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Data source type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#type KendraDataSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#custom_document_enrichment_configuration KendraDataSource#custom_document_enrichment_configuration}.
	CustomDocumentEnrichmentConfiguration *KendraDataSourceCustomDocumentEnrichmentConfiguration `field:"optional" json:"customDocumentEnrichmentConfiguration" yaml:"customDocumentEnrichmentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#data_source_configuration KendraDataSource#data_source_configuration}.
	DataSourceConfiguration *KendraDataSourceDataSourceConfiguration `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// Description of data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#description KendraDataSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The code for a language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#language_code KendraDataSource#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#role_arn KendraDataSource#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#schedule KendraDataSource#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Tags for labeling the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#tags KendraDataSource#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


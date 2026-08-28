// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataSourceConfig struct {
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
	// Specifies a raw data source location to ingest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#data_source_configuration BedrockDataSource#data_source_configuration}
	DataSourceConfiguration *BedrockDataSourceDataSourceConfiguration `field:"required" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// The unique identifier of the knowledge base to which to add the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#knowledge_base_id BedrockDataSource#knowledge_base_id}
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The name of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#name BedrockDataSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The deletion policy for the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#data_deletion_policy BedrockDataSource#data_deletion_policy}
	DataDeletionPolicy *string `field:"optional" json:"dataDeletionPolicy" yaml:"dataDeletionPolicy"`
	// Description of the Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#description BedrockDataSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Contains details about the server-side encryption for the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#server_side_encryption_configuration BedrockDataSource#server_side_encryption_configuration}
	ServerSideEncryptionConfiguration *BedrockDataSourceServerSideEncryptionConfiguration `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Details about how to chunk the documents in the data source.
	//
	// A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#vector_ingestion_configuration BedrockDataSource#vector_ingestion_configuration}
	VectorIngestionConfiguration *BedrockDataSourceVectorIngestionConfiguration `field:"optional" json:"vectorIngestionConfiguration" yaml:"vectorIngestionConfiguration"`
}


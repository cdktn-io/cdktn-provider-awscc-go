// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationStorageConfigurationsAwsDataCatalogConfiguration struct {
	// List of table names in AWS Data Catalog. Must follow two part notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#table_names BedrockKnowledgeBase#table_names}
	TableNames *[]*string `field:"optional" json:"tableNames" yaml:"tableNames"`
}


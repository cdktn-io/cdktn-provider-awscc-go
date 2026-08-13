// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfiguration struct {
	// Contains the storage configuration of the knowledge base in MongoDb Atlas Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#mongo_db_atlas_configuration BedrockKnowledgeBase#mongo_db_atlas_configuration}
	MongoDbAtlasConfiguration *BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfiguration `field:"optional" json:"mongoDbAtlasConfiguration" yaml:"mongoDbAtlasConfiguration"`
	// Contains the configurations to use Neptune Analytics as Vector Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#neptune_analytics_configuration BedrockKnowledgeBase#neptune_analytics_configuration}
	NeptuneAnalyticsConfiguration *BedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfiguration `field:"optional" json:"neptuneAnalyticsConfiguration" yaml:"neptuneAnalyticsConfiguration"`
	// Contains the storage configuration of the knowledge base in Amazon OpenSearch Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#opensearch_managed_cluster_configuration BedrockKnowledgeBase#opensearch_managed_cluster_configuration}
	OpensearchManagedClusterConfiguration *BedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfiguration `field:"optional" json:"opensearchManagedClusterConfiguration" yaml:"opensearchManagedClusterConfiguration"`
	// Contains the storage configuration of the knowledge base in Amazon OpenSearch Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#opensearch_serverless_configuration BedrockKnowledgeBase#opensearch_serverless_configuration}
	OpensearchServerlessConfiguration *BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfiguration `field:"optional" json:"opensearchServerlessConfiguration" yaml:"opensearchServerlessConfiguration"`
	// Contains the storage configuration of the knowledge base in Pinecone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#pinecone_configuration BedrockKnowledgeBase#pinecone_configuration}
	PineconeConfiguration *BedrockKnowledgeBaseStorageConfigurationPineconeConfiguration `field:"optional" json:"pineconeConfiguration" yaml:"pineconeConfiguration"`
	// Contains details about the storage configuration of the knowledge base in Amazon RDS.
	//
	// For more information, see Create a vector index in Amazon RDS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#rds_configuration BedrockKnowledgeBase#rds_configuration}
	RdsConfiguration *BedrockKnowledgeBaseStorageConfigurationRdsConfiguration `field:"optional" json:"rdsConfiguration" yaml:"rdsConfiguration"`
	// Contains the storage configuration of the knowledge base for S3 vectors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#s3_vectors_configuration BedrockKnowledgeBase#s3_vectors_configuration}
	S3VectorsConfiguration *BedrockKnowledgeBaseStorageConfigurationS3VectorsConfiguration `field:"optional" json:"s3VectorsConfiguration" yaml:"s3VectorsConfiguration"`
	// The storage type of a knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#type BedrockKnowledgeBase#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


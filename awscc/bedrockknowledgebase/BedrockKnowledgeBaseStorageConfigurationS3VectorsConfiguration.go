// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationS3VectorsConfiguration struct {
	// The Amazon Resource Name (ARN) of the vector index used for the knowledge base.
	//
	// This ARN identifies the specific vector index resource within Amazon Bedrock.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#index_arn BedrockKnowledgeBase#index_arn}
	IndexArn *string `field:"optional" json:"indexArn" yaml:"indexArn"`
	// The name of the vector index used for the knowledge base.
	//
	// This name identifies the vector index within the Amazon Bedrock service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#index_name BedrockKnowledgeBase#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The Amazon Resource Name (ARN) of the S3 bucket where vector embeddings are stored.
	//
	// This bucket contains the vector data used by the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#vector_bucket_arn BedrockKnowledgeBase#vector_bucket_arn}
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
}


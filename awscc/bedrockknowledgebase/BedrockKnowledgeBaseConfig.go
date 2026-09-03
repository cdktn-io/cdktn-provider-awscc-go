// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockKnowledgeBaseConfig struct {
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
	// Contains details about the embeddings model used for the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#knowledge_base_configuration BedrockKnowledgeBase#knowledge_base_configuration}
	KnowledgeBaseConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfiguration `field:"required" json:"knowledgeBaseConfiguration" yaml:"knowledgeBaseConfiguration"`
	// The name of the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#name BedrockKnowledgeBase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the IAM role with permissions to invoke API operations on the knowledge base.
	//
	// The ARN must begin with AmazonBedrockExecutionRoleForKnowledgeBase_
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#role_arn BedrockKnowledgeBase#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Description of the Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#description BedrockKnowledgeBase#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The vector store service in which the knowledge base is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#storage_configuration BedrockKnowledgeBase#storage_configuration}
	StorageConfiguration *BedrockKnowledgeBaseStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#tags BedrockKnowledgeBase#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


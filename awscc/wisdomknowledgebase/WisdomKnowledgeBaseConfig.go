// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomknowledgebase

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomKnowledgeBaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#knowledge_base_type WisdomKnowledgeBase#knowledge_base_type}.
	KnowledgeBaseType *string `field:"required" json:"knowledgeBaseType" yaml:"knowledgeBaseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#name WisdomKnowledgeBase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#description WisdomKnowledgeBase#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#rendering_configuration WisdomKnowledgeBase#rendering_configuration}.
	RenderingConfiguration *WisdomKnowledgeBaseRenderingConfiguration `field:"optional" json:"renderingConfiguration" yaml:"renderingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#server_side_encryption_configuration WisdomKnowledgeBase#server_side_encryption_configuration}.
	ServerSideEncryptionConfiguration *WisdomKnowledgeBaseServerSideEncryptionConfiguration `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#source_configuration WisdomKnowledgeBase#source_configuration}.
	SourceConfiguration *WisdomKnowledgeBaseSourceConfiguration `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#tags WisdomKnowledgeBase#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#vector_ingestion_configuration WisdomKnowledgeBase#vector_ingestion_configuration}.
	VectorIngestionConfiguration *WisdomKnowledgeBaseVectorIngestionConfiguration `field:"optional" json:"vectorIngestionConfiguration" yaml:"vectorIngestionConfiguration"`
}


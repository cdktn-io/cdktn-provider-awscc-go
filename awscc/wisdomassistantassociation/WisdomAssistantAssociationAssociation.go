// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomassistantassociation


type WisdomAssistantAssociationAssociation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_assistant_association#external_bedrock_knowledge_base_config WisdomAssistantAssociation#external_bedrock_knowledge_base_config}.
	ExternalBedrockKnowledgeBaseConfig *WisdomAssistantAssociationAssociationExternalBedrockKnowledgeBaseConfig `field:"optional" json:"externalBedrockKnowledgeBaseConfig" yaml:"externalBedrockKnowledgeBaseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_assistant_association#knowledge_base_id WisdomAssistantAssociation#knowledge_base_id}.
	KnowledgeBaseId *string `field:"optional" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
}


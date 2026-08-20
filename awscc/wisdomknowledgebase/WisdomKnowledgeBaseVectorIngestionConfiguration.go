// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomknowledgebase


type WisdomKnowledgeBaseVectorIngestionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#chunking_configuration WisdomKnowledgeBase#chunking_configuration}.
	ChunkingConfiguration *WisdomKnowledgeBaseVectorIngestionConfigurationChunkingConfiguration `field:"optional" json:"chunkingConfiguration" yaml:"chunkingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_knowledge_base#parsing_configuration WisdomKnowledgeBase#parsing_configuration}.
	ParsingConfiguration *WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfiguration `field:"optional" json:"parsingConfiguration" yaml:"parsingConfiguration"`
}


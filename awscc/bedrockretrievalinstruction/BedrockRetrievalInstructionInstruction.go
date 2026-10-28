// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockretrievalinstruction


type BedrockRetrievalInstructionInstruction struct {
	// How the answer is structured, including how sources are cited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrock_retrieval_instruction#answer_format BedrockRetrievalInstruction#answer_format}
	AnswerFormat *string `field:"optional" json:"answerFormat" yaml:"answerFormat"`
	// Whether and how to answer: how strongly an answer must be supported, and what to do with questions outside scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrock_retrieval_instruction#answer_policy BedrockRetrievalInstruction#answer_policy}
	AnswerPolicy *string `field:"optional" json:"answerPolicy" yaml:"answerPolicy"`
	// Who the assistant is: identity, tone and language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrock_retrieval_instruction#persona BedrockRetrievalInstruction#persona}
	Persona *string `field:"optional" json:"persona" yaml:"persona"`
	// How to interpret and decompose the question, and what to prioritise when searching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrock_retrieval_instruction#retrieval_guidance BedrockRetrievalInstruction#retrieval_guidance}
	RetrievalGuidance *string `field:"optional" json:"retrievalGuidance" yaml:"retrievalGuidance"`
}


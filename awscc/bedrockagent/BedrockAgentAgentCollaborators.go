// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentAgentCollaborators struct {
	// Agent descriptor for agent collaborator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_agent#agent_descriptor BedrockAgent#agent_descriptor}
	AgentDescriptor *BedrockAgentAgentCollaboratorsAgentDescriptor `field:"optional" json:"agentDescriptor" yaml:"agentDescriptor"`
	// Agent collaborator instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_agent#collaboration_instruction BedrockAgent#collaboration_instruction}
	CollaborationInstruction *string `field:"optional" json:"collaborationInstruction" yaml:"collaborationInstruction"`
	// Agent collaborator name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_agent#collaborator_name BedrockAgent#collaborator_name}
	CollaboratorName *string `field:"optional" json:"collaboratorName" yaml:"collaboratorName"`
	// Relay conversation history state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_agent#relay_conversation_history BedrockAgent#relay_conversation_history}
	RelayConversationHistory *string `field:"optional" json:"relayConversationHistory" yaml:"relayConversationHistory"`
}


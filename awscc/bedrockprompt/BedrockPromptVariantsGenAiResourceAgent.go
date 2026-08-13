// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsGenAiResourceAgent struct {
	// Arn representation of the Agent Alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_prompt#agent_identifier BedrockPrompt#agent_identifier}
	AgentIdentifier *string `field:"optional" json:"agentIdentifier" yaml:"agentIdentifier"`
}


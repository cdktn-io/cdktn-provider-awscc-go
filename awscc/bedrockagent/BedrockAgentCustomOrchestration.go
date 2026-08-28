// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentCustomOrchestration struct {
	// Types of executors for custom orchestration strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_agent#executor BedrockAgent#executor}
	Executor *BedrockAgentCustomOrchestrationExecutor `field:"optional" json:"executor" yaml:"executor"`
}


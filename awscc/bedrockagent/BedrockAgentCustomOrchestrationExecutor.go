// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentCustomOrchestrationExecutor struct {
	// ARN of a Lambda.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_agent#lambda BedrockAgent#lambda}
	Lambda *string `field:"optional" json:"lambda" yaml:"lambda"`
}


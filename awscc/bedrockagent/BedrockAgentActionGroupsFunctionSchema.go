// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentActionGroupsFunctionSchema struct {
	// List of Function definitions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent#functions BedrockAgent#functions}
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
}


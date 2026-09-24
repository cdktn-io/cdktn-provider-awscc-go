// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentActionGroupsFunctionSchemaFunctions struct {
	// Description of function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_agent#description BedrockAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_agent#name BedrockAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map of parameter name and detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_agent#parameters BedrockAgent#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// ENUM to check if action requires user confirmation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_agent#require_confirmation BedrockAgent#require_confirmation}
	RequireConfirmation *string `field:"optional" json:"requireConfirmation" yaml:"requireConfirmation"`
}


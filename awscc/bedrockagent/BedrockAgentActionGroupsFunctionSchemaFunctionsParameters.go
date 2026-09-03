// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentActionGroupsFunctionSchemaFunctionsParameters struct {
	// Description of function parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_agent#description BedrockAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about if a parameter is required for function call. Default to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_agent#required BedrockAgent#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Parameter Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_agent#type BedrockAgent#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


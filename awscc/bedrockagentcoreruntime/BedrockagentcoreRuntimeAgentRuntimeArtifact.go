// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAgentRuntimeArtifact struct {
	// Representation of a code configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#code_configuration BedrockagentcoreRuntime#code_configuration}
	CodeConfiguration *BedrockagentcoreRuntimeAgentRuntimeArtifactCodeConfiguration `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#container_configuration BedrockagentcoreRuntime#container_configuration}.
	ContainerConfiguration *BedrockagentcoreRuntimeAgentRuntimeArtifactContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAgentRuntimeArtifactCodeConfiguration struct {
	// Object represents source code from zip file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#code BedrockagentcoreRuntime#code}
	Code *BedrockagentcoreRuntimeAgentRuntimeArtifactCodeConfigurationCode `field:"optional" json:"code" yaml:"code"`
	// List of entry points.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#entry_point BedrockagentcoreRuntime#entry_point}
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Managed runtime types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#runtime BedrockagentcoreRuntime#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}


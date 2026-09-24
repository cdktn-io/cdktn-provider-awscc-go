// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryNamespaceKeys struct {
	// A namespace variable key name.
	//
	// Must start with a lowercase letter and contain only lowercase alphanumeric characters. Cannot be a built-in variable (memoryStrategyId, sessionId, actorId).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#key BedrockagentcoreMemory#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Validation rules for namespace variable values. Multiple rules can be specified and all must pass.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#validation BedrockagentcoreMemory#validation}
	Validation *BedrockagentcoreMemoryNamespaceKeysValidation `field:"optional" json:"validation" yaml:"validation"`
}


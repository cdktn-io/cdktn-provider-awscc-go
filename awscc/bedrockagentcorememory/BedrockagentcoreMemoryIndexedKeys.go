// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryIndexedKeys struct {
	// Key name for metadata fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_memory#key BedrockagentcoreMemory#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Supported data types for metadata values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_memory#type BedrockagentcoreMemory#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


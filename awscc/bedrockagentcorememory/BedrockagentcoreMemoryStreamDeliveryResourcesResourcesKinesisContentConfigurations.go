// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryStreamDeliveryResourcesResourcesKinesisContentConfigurations struct {
	// The level of content detail to deliver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#level BedrockagentcoreMemory#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// The type of content to deliver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#type BedrockagentcoreMemory#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


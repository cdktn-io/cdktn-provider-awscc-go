// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockblueprint


type BedrockBlueprintTags struct {
	// Key for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_blueprint#key BedrockBlueprint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_blueprint#value BedrockBlueprint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


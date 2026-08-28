// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueregistry


type GlueRegistryTags struct {
	// A key to identify the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_registry#key GlueRegistry#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponding tag value for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_registry#value GlueRegistry#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


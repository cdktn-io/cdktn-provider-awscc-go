// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrrepositorycreationtemplate


type EcrRepositoryCreationTemplateResourceTags struct {
	// One part of a key-value pair that make up a tag.
	//
	// A ``key`` is a general label that acts like a category for more specific tag values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecr_repository_creation_template#key EcrRepositoryCreationTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A ``value`` acts as a descriptor within a tag category (key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecr_repository_creation_template#value EcrRepositoryCreationTemplate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


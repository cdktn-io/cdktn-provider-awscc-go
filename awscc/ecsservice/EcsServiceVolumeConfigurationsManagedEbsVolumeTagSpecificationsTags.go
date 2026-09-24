// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceVolumeConfigurationsManagedEbsVolumeTagSpecificationsTags struct {
	// One part of a key-value pair that make up a tag.
	//
	// A ``key`` is a general label that acts like a category for more specific tag values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service#key EcsService#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The optional part of a key-value pair that make up a tag.
	//
	// A ``value`` acts as a descriptor within a tag category (key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service#value EcsService#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterTags struct {
	// One part of a key-value pair that make up a tag.
	//
	// A ``key`` is a general label that acts like a category for more specific tag values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_cluster#key EcsCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The optional part of a key-value pair that make up a tag.
	//
	// A ``value`` acts as a descriptor within a tag category (key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_cluster#value EcsCluster#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


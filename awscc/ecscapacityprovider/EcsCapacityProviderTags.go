// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#key EcsCapacityProvider#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#value EcsCapacityProvider#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


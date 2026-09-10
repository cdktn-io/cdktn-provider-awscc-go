// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsclustercapacityproviderassociations


type EcsClusterCapacityProviderAssociationsDefaultCapacityProviderStrategy struct {
	// If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_cluster_capacity_provider_associations#capacity_provider EcsClusterCapacityProviderAssociations#capacity_provider}
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_cluster_capacity_provider_associations#base EcsClusterCapacityProviderAssociations#base}.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_cluster_capacity_provider_associations#weight EcsClusterCapacityProviderAssociations#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceVolumeConfigurations struct {
	// The configuration for the Amazon EBS volume that Amazon ECS creates and manages on your behalf.
	//
	// These settings are used to create each Amazon EBS volume, with one volume created for each task in the service. The Amazon EBS volumes are visible in your account in the Amazon EC2 console once they are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#managed_ebs_volume EcsService#managed_ebs_volume}
	ManagedEbsVolume *EcsServiceVolumeConfigurationsManagedEbsVolume `field:"optional" json:"managedEbsVolume" yaml:"managedEbsVolume"`
	// The name of the volume.
	//
	// This value must match the volume name from the ``Volume`` object in the task definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#name EcsService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


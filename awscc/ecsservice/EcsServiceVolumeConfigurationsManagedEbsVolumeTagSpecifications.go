// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceVolumeConfigurationsManagedEbsVolumeTagSpecifications struct {
	// Determines whether to propagate the tags from the task definition to the Amazon EBS volume.
	//
	// Tags can only propagate to a ``SERVICE`` specified in ``ServiceVolumeConfiguration``. If no value is specified, the tags aren't propagated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#propagate_tags EcsService#propagate_tags}
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The type of volume resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#resource_type EcsService#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags applied to this Amazon EBS volume. ``AmazonECSCreated`` and ``AmazonECSManaged`` are reserved tags that can't be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#tags EcsService#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumesEfsVolumeConfigurationAuthorizationConfig struct {
	// The Amazon EFS access point ID to use.
	//
	// If an access point is specified, the root directory value specified in the ``EFSVolumeConfiguration`` must either be omitted or set to ``/`` which will enforce the path set on the EFS access point. If an access point is used, transit encryption must be on in the ``EFSVolumeConfiguration``. For more information, see [Working with Amazon EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) in the *Amazon Elastic File System User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#access_point_id EcsTaskDefinition#access_point_id}
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Determines whether to use the Amazon ECS task role defined in a task definition when mounting the Amazon EFS file system.
	//
	// If it is turned on, transit encryption must be turned on in the ``EFSVolumeConfiguration``. If this parameter is omitted, the default value of ``DISABLED`` is used. For more information, see [Using Amazon EFS access points](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html#efs-volume-accesspoints) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#iam EcsTaskDefinition#iam}
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}


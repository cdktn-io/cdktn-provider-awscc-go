// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisetask


type IotsitewiseTaskTaskConfigurationContainerTaskConfigurationMountsSourceS3AccessPoint struct {
	// The Amazon Resource Name (ARN) of the Amazon S3 access point.
	//
	// The mount reads objects from the bucket associated with this access point. Access is governed by the access point policy and the task execution role's IAM permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_task#access_point_arn IotsitewiseTask#access_point_arn}
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// An object key name prefix.
	//
	// If specified, the mount includes only objects whose keys begin with this prefix. To include all objects at the access point, omit this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_task#prefix IotsitewiseTask#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}


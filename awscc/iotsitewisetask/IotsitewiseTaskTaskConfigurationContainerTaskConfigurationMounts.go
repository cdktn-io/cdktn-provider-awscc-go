// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisetask


type IotsitewiseTaskTaskConfigurationContainerTaskConfigurationMounts struct {
	// A unique name for the mount within the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_task#name IotsitewiseTask#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The relative path under the service-owned mount root where this mount is attached inside the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_task#relative_path IotsitewiseTask#relative_path}
	RelativePath *string `field:"optional" json:"relativePath" yaml:"relativePath"`
	// The data source configuration for a mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_task#source IotsitewiseTask#source}
	Source *IotsitewiseTaskTaskConfigurationContainerTaskConfigurationMountsSource `field:"optional" json:"source" yaml:"source"`
	// The type of storage used for the mount inside the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_task#storage_type IotsitewiseTask#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
}


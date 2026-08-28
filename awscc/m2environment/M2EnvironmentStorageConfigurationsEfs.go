// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package m2environment


type M2EnvironmentStorageConfigurationsEfs struct {
	// The file system identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/m2_environment#file_system_id M2Environment#file_system_id}
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The mount point for the file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/m2_environment#mount_point M2Environment#mount_point}
	MountPoint *string `field:"optional" json:"mountPoint" yaml:"mountPoint"`
}


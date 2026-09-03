// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsaccesspoint


type EfsAccessPointRootDirectory struct {
	// (Optional) Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory``.
	//
	// If the ``RootDirectory`` > ``Path`` specified does not exist, EFS creates the root directory using the ``CreationInfo`` settings when a client connects to an access point. When specifying the ``CreationInfo``, you must provide values for all properties.
	//   If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` > ``Path`` does not exist, attempts to mount the file system using the access point will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/efs_access_point#creation_info EfsAccessPoint#creation_info}
	CreationInfo *EfsAccessPointRootDirectoryCreationInfo `field:"optional" json:"creationInfo" yaml:"creationInfo"`
	// Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system.
	//
	// A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the ``CreationInfo``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/efs_access_point#path EfsAccessPoint#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}


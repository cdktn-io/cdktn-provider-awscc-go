// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsfilesystem


type EfsFileSystemBackupPolicy struct {
	// Set the backup policy status for the file system.
	//
	// +  *ENABLED* - Turns automatic backups on for the file system.
	//   +  *DISABLED* - Turns automatic backups off for the file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/efs_file_system#status EfsFileSystem#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


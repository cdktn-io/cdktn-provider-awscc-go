// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsfilesystem


type EfsFileSystemFileSystemProtection struct {
	// The status of the file system's replication overwrite protection.
	//
	// +  ``ENABLED`` ? The file system cannot be used as the destination file system in a replication configuration. The file system is writeable. Replication overwrite protection is ``ENABLED`` by default.
	//   +  ``DISABLED`` ? The file system can be used as the destination file system in a replication configuration. The file system is read-only and can only be modified by EFS replication.
	//   +  ``REPLICATING`` ? The file system is being used as the destination file system in a replication configuration. The file system is read-only and is modified only by EFS replication.
	//
	//  If the replication configuration is deleted, the file system's replication overwrite protection is re-enabled, the file system becomes writeable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/efs_file_system#replication_overwrite_protection EfsFileSystem#replication_overwrite_protection}
	ReplicationOverwriteProtection *string `field:"optional" json:"replicationOverwriteProtection" yaml:"replicationOverwriteProtection"`
}


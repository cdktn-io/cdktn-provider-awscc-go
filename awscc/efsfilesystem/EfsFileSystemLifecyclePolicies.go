// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsfilesystem


type EfsFileSystemLifecyclePolicies struct {
	// The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Archive storage.
	//
	// Metadata operations such as listing the contents of a directory don't count as file access events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/efs_file_system#transition_to_archive EfsFileSystem#transition_to_archive}
	TransitionToArchive *string `field:"optional" json:"transitionToArchive" yaml:"transitionToArchive"`
	// The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Infrequent Access (IA) storage.
	//
	// Metadata operations such as listing the contents of a directory don't count as file access events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/efs_file_system#transition_to_ia EfsFileSystem#transition_to_ia}
	TransitionToIa *string `field:"optional" json:"transitionToIa" yaml:"transitionToIa"`
	// Whether to move files back to primary (Standard) storage after they are accessed in IA or Archive storage.
	//
	// Metadata operations such as listing the contents of a directory don't count as file access events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/efs_file_system#transition_to_primary_storage_class EfsFileSystem#transition_to_primary_storage_class}
	TransitionToPrimaryStorageClass *string `field:"optional" json:"transitionToPrimaryStorageClass" yaml:"transitionToPrimaryStorageClass"`
}


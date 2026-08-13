// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsfilesystem


type EfsFileSystemFileSystemTags struct {
	// The tag key (String). The key can't start with ``aws:``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/efs_file_system#key EfsFileSystem#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/efs_file_system#value EfsFileSystem#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


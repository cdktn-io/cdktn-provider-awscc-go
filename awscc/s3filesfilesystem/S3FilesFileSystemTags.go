// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesfilesystem


type S3FilesFileSystemTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3files_file_system#key S3FilesFileSystem#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3files_file_system#value S3FilesFileSystem#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


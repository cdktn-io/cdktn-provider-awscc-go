// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesfilesystem


type S3FilesFileSystemSynchronizationConfigurationImportDataRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3files_file_system#prefix S3FilesFileSystem#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3files_file_system#size_less_than S3FilesFileSystem#size_less_than}.
	SizeLessThan *float64 `field:"optional" json:"sizeLessThan" yaml:"sizeLessThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3files_file_system#trigger S3FilesFileSystem#trigger}.
	Trigger *string `field:"optional" json:"trigger" yaml:"trigger"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesfilesystem


type S3FilesFileSystemSynchronizationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3files_file_system#expiration_data_rules S3FilesFileSystem#expiration_data_rules}.
	ExpirationDataRules interface{} `field:"optional" json:"expirationDataRules" yaml:"expirationDataRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3files_file_system#import_data_rules S3FilesFileSystem#import_data_rules}.
	ImportDataRules interface{} `field:"optional" json:"importDataRules" yaml:"importDataRules"`
}


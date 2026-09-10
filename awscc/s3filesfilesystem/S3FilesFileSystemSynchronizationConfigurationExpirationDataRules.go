// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesfilesystem


type S3FilesFileSystemSynchronizationConfigurationExpirationDataRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3files_file_system#days_after_last_access S3FilesFileSystem#days_after_last_access}.
	DaysAfterLastAccess *float64 `field:"optional" json:"daysAfterLastAccess" yaml:"daysAfterLastAccess"`
}


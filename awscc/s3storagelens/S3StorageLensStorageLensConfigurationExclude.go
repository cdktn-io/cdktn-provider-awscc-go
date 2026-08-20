// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationExclude struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#buckets S3StorageLens#buckets}.
	Buckets *[]*string `field:"optional" json:"buckets" yaml:"buckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#regions S3StorageLens#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


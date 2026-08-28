// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryption struct {
	// AWS KMS server-side encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#ssekms S3StorageLens#ssekms}
	Ssekms *S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionSsekms `field:"optional" json:"ssekms" yaml:"ssekms"`
	// S3 default server-side encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#sses3 S3StorageLens#sses3}
	Sses3 *string `field:"optional" json:"sses3" yaml:"sses3"`
}


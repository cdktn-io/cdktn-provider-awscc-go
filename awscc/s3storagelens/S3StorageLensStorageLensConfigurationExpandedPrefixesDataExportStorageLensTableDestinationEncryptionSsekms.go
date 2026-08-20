// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionSsekms struct {
	// The ARN of the KMS key to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#key_id S3StorageLens#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}


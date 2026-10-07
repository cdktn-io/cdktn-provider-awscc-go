// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketMetadataConfigurationJournalTableConfiguration struct {
	// The encryption configuration for the journal table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#encryption_configuration S3Bucket#encryption_configuration}
	EncryptionConfiguration *S3BucketMetadataConfigurationJournalTableConfigurationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The journal table record expiration settings for the journal table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#record_expiration S3Bucket#record_expiration}
	RecordExpiration *S3BucketMetadataConfigurationJournalTableConfigurationRecordExpiration `field:"optional" json:"recordExpiration" yaml:"recordExpiration"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketMetadataConfigurationJournalTableConfigurationRecordExpiration struct {
	// If you enable journal table record expiration, you can set the number of days to retain your journal table records.
	//
	// Journal table records must be retained for a minimum of 7 days. To set this value, specify any whole number from ``7`` to ``2147483647``. For example, to retain your journal table records for one year, set this value to ``365``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#days S3Bucket#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Specifies whether journal table record expiration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#expiration S3Bucket#expiration}
	Expiration *string `field:"optional" json:"expiration" yaml:"expiration"`
}


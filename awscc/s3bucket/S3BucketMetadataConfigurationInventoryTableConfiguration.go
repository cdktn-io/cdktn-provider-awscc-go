// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketMetadataConfigurationInventoryTableConfiguration struct {
	// The configuration state of the inventory table, indicating whether the inventory table is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#configuration_state S3Bucket#configuration_state}
	ConfigurationState *string `field:"optional" json:"configurationState" yaml:"configurationState"`
	// The encryption configuration for the inventory table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#encryption_configuration S3Bucket#encryption_configuration}
	EncryptionConfiguration *S3BucketMetadataConfigurationInventoryTableConfigurationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}


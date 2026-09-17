// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#kms_encryption_config KinesisfirehoseDeliveryStream#kms_encryption_config}.
	KmsEncryptionConfig *KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfig `field:"optional" json:"kmsEncryptionConfig" yaml:"kmsEncryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#no_encryption_config KinesisfirehoseDeliveryStream#no_encryption_config}.
	NoEncryptionConfig *string `field:"optional" json:"noEncryptionConfig" yaml:"noEncryptionConfig"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisstream


type KinesisStreamStreamEncryption struct {
	// The encryption type to use. The only valid value is KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesis_stream#encryption_type KinesisStream#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The GUID for the customer-managed AWS KMS key to use for encryption.
	//
	// This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesis_stream#key_id KinesisStream#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}


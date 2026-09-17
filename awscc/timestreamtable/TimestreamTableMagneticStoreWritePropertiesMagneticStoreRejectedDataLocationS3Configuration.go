// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamtable


type TimestreamTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3Configuration struct {
	// The bucket name used to store the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_table#bucket_name TimestreamTable#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Either SSE_KMS or SSE_S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_table#encryption_option TimestreamTable#encryption_option}
	EncryptionOption *string `field:"optional" json:"encryptionOption" yaml:"encryptionOption"`
	// Must be provided if SSE_KMS is specified as the encryption option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_table#kms_key_id TimestreamTable#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// String used to prefix all data in the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_table#object_key_prefix TimestreamTable#object_key_prefix}
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}


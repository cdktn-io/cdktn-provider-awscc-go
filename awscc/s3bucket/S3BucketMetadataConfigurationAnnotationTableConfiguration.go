// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketMetadataConfigurationAnnotationTableConfiguration struct {
	// Specifies whether the annotation table configuration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#configuration_state S3Bucket#configuration_state}
	ConfigurationState *string `field:"optional" json:"configurationState" yaml:"configurationState"`
	// The encryption configuration for the annotation table.
	//
	// To encrypt your annotation table with server-side encryption using AWS Key Management Service (AWS KMS) keys (SSE-KMS), set ``SseAlgorithm`` to ``aws:kms``. You must also set ``KmsKeyArn`` to the ARN of a customer managed KMS key in the same Region where your general purpose bucket is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#encryption_configuration S3Bucket#encryption_configuration}
	EncryptionConfiguration *S3BucketMetadataConfigurationAnnotationTableConfigurationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The ARN of the IAM role that grants Amazon S3 Metadata permission to read annotations from your bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#role S3Bucket#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}


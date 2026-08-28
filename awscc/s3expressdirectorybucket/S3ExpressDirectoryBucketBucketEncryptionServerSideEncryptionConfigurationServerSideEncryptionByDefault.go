// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressdirectorybucket


type S3ExpressDirectoryBucketBucketEncryptionServerSideEncryptionConfigurationServerSideEncryptionByDefault struct {
	// AWS Key Management Service (KMS) customer managed key ID to use for the default encryption.
	//
	// This parameter is allowed only if SSEAlgorithm is set to aws:kms. You can specify this parameter with the key ID or the Amazon Resource Name (ARN) of the KMS key
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3express_directory_bucket#kms_master_key_id S3ExpressDirectoryBucket#kms_master_key_id}
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3express_directory_bucket#sse_algorithm S3ExpressDirectoryBucket#sse_algorithm}.
	SseAlgorithm *string `field:"optional" json:"sseAlgorithm" yaml:"sseAlgorithm"`
}


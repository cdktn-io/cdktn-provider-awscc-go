// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3vectorsvectorbucket


type S3VectorsVectorBucketEncryptionConfiguration struct {
	// AWS Key Management Service (KMS) customer managed key ID to use for the encryption configuration.
	//
	// This parameter is allowed if and only if sseType is set to aws:kms
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3vectors_vector_bucket#kms_key_arn S3VectorsVectorBucket#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The server-side encryption type to use for the encryption configuration of the vector bucket.
	//
	// By default, if you don't specify, all new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3 managed keys (SSE-S3), specifically AES256.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3vectors_vector_bucket#sse_type S3VectorsVectorBucket#sse_type}
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}


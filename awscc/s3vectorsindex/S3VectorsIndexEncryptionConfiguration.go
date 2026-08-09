// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3vectorsindex


type S3VectorsIndexEncryptionConfiguration struct {
	// AWS Key Management Service (KMS) customer managed key ID to use for the encryption configuration.
	//
	// This parameter is allowed if and only if sseType is set to aws:kms
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3vectors_index#kms_key_arn S3VectorsIndex#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Defines the server-side encryption type for index encryption configuration. Defaults to the parent vector bucket's encryption settings when unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3vectors_index#sse_type S3VectorsIndex#sse_type}
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}


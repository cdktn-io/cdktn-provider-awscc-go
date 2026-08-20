// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmins3tableintegration


type ObservabilityadminS3TableIntegrationEncryption struct {
	// The server-side encryption algorithm used to encrypt the S3 Table(s) data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_s3_table_integration#sse_algorithm ObservabilityadminS3TableIntegration#sse_algorithm}
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// The ARN of the KMS key used to encrypt the S3 Table Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_s3_table_integration#kms_key_arn ObservabilityadminS3TableIntegration#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}


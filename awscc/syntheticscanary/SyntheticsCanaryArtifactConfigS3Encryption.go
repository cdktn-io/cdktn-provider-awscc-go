// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryArtifactConfigS3Encryption struct {
	// Encryption mode for encrypting artifacts when uploading to S3. Valid values: SSE_S3 and SSE_KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#encryption_mode SyntheticsCanary#encryption_mode}
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// KMS key Arn for encrypting artifacts when uploading to S3.
	//
	// You must specify KMS key Arn for SSE_KMS encryption mode only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#kms_key_arn SyntheticsCanary#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}


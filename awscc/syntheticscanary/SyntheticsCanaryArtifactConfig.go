// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryArtifactConfig struct {
	// Encryption configuration for uploading artifacts to S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/synthetics_canary#s3_encryption SyntheticsCanary#s3_encryption}
	S3Encryption *SyntheticsCanaryArtifactConfigS3Encryption `field:"optional" json:"s3Encryption" yaml:"s3Encryption"`
}


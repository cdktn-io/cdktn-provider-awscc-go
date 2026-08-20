// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineEncryptionAtRestOptions struct {
	// The KMS key to use for encrypting data. By default an AWS owned key is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/osis_pipeline#kms_key_arn OsisPipeline#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}


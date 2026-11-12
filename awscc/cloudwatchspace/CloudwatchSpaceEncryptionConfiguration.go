// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchspace


type CloudwatchSpaceEncryptionConfiguration struct {
	// Which kind of key to use. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_space#encryption_strategy CloudwatchSpace#encryption_strategy}
	EncryptionStrategy *string `field:"optional" json:"encryptionStrategy" yaml:"encryptionStrategy"`
	// Customer managed KMS key ARN.
	//
	// Required when encryptionStrategy is CUSTOMER_MANAGED, and must be omitted when it is AWS_OWNED. Must be a symmetric ENCRYPT_DECRYPT key in the caller's account and region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_space#kms_key_arn CloudwatchSpace#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}


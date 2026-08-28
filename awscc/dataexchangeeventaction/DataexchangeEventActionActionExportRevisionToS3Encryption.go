// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionActionExportRevisionToS3Encryption struct {
	// The Amazon Resource Name (ARN) of the AWS KMS key you want to use to encrypt the Amazon S3 objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dataexchange_event_action#kms_key_arn DataexchangeEventAction#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The type of server side encryption used for encrypting the objects in Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dataexchange_event_action#type DataexchangeEventAction#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


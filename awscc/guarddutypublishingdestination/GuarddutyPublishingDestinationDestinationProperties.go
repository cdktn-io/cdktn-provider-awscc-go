// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutypublishingdestination


type GuarddutyPublishingDestinationDestinationProperties struct {
	// The ARN of the resource to publish to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_publishing_destination#destination_arn GuarddutyPublishingDestination#destination_arn}
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// The ARN of the KMS key to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_publishing_destination#kms_key_arn GuarddutyPublishingDestination#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}


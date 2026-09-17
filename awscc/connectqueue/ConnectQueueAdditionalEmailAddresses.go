// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectqueue


type ConnectQueueAdditionalEmailAddresses struct {
	// The Amazon Resource Name (ARN) of the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_queue#email_address_arn ConnectQueue#email_address_arn}
	EmailAddressArn *string `field:"optional" json:"emailAddressArn" yaml:"emailAddressArn"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectemailaddress


type ConnectEmailAddressAliasConfigurations struct {
	// The identifier of the email address alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_email_address#email_address_arn ConnectEmailAddress#email_address_arn}
	EmailAddressArn *string `field:"optional" json:"emailAddressArn" yaml:"emailAddressArn"`
}


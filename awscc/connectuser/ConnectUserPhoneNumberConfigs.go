// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserPhoneNumberConfigs struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_user#channel ConnectUser#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The phone number for the user's desk phone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_user#phone_number ConnectUser#phone_number}
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The phone type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_user#phone_type ConnectUser#phone_type}
	PhoneType *string `field:"optional" json:"phoneType" yaml:"phoneType"`
}


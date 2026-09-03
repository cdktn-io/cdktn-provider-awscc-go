// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoicephonenumber


type SmsvoicePhoneNumberTwoWay struct {
	// The Amazon Resource Name (ARN) of the two way channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_phone_number#channel_arn SmsvoicePhoneNumber#channel_arn}
	ChannelArn *string `field:"optional" json:"channelArn" yaml:"channelArn"`
	// An optional IAM Role Arn for a service to assume, to be able to post inbound SMS messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_phone_number#channel_role SmsvoicePhoneNumber#channel_role}
	ChannelRole *string `field:"optional" json:"channelRole" yaml:"channelRole"`
	// By default this is set to false.
	//
	// When set to true you can receive incoming text messages from your end recipients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_phone_number#enabled SmsvoicePhoneNumber#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}


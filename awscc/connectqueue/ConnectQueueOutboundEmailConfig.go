// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectqueue


type ConnectQueueOutboundEmailConfig struct {
	// The email address connect resource ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_queue#outbound_email_address_id ConnectQueue#outbound_email_address_id}
	OutboundEmailAddressId *string `field:"optional" json:"outboundEmailAddressId" yaml:"outboundEmailAddressId"`
}


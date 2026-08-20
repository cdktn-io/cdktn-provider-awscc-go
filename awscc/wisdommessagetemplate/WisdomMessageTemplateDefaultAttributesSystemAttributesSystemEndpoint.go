// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateDefaultAttributesSystemAttributesSystemEndpoint struct {
	// The customer's phone number if used with customerEndpoint, or the number the customer dialed to call your contact center if used with systemEndpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_message_template#address WisdomMessageTemplate#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
}


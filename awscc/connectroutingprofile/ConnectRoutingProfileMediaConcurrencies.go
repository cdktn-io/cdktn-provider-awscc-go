// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectroutingprofile


type ConnectRoutingProfileMediaConcurrencies struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#channel ConnectRoutingProfile#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The number of contacts an agent can have on a channel simultaneously.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#concurrency ConnectRoutingProfile#concurrency}
	Concurrency *float64 `field:"required" json:"concurrency" yaml:"concurrency"`
	// Defines the cross-channel routing behavior that allows an agent working on a contact in one channel to be offered a contact from a different channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#cross_channel_behavior ConnectRoutingProfile#cross_channel_behavior}
	CrossChannelBehavior *ConnectRoutingProfileMediaConcurrenciesCrossChannelBehavior `field:"optional" json:"crossChannelBehavior" yaml:"crossChannelBehavior"`
}


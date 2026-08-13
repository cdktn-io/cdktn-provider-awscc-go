// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsendpoint


type EventsEndpointRoutingConfigFailoverConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/events_endpoint#primary EventsEndpoint#primary}.
	Primary *EventsEndpointRoutingConfigFailoverConfigPrimary `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/events_endpoint#secondary EventsEndpoint#secondary}.
	Secondary *EventsEndpointRoutingConfigFailoverConfigSecondary `field:"required" json:"secondary" yaml:"secondary"`
}


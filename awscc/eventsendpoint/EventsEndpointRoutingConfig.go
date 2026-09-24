// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsendpoint


type EventsEndpointRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/events_endpoint#failover_config EventsEndpoint#failover_config}.
	FailoverConfig *EventsEndpointRoutingConfigFailoverConfig `field:"required" json:"failoverConfig" yaml:"failoverConfig"`
}


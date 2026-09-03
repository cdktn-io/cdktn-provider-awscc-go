// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsendpoint


type EventsEndpointRoutingConfigFailoverConfigPrimary struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_endpoint#health_check EventsEndpoint#health_check}.
	HealthCheck *string `field:"required" json:"healthCheck" yaml:"healthCheck"`
}


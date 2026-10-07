// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventsource


type Eventsv2EventSourceConfiguration struct {
	// Configuration for forwarding a single AWS service's events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_source#aws_service_events_configuration Eventsv2EventSource#aws_service_events_configuration}
	AwsServiceEventsConfiguration *Eventsv2EventSourceConfigurationAwsServiceEventsConfiguration `field:"optional" json:"awsServiceEventsConfiguration" yaml:"awsServiceEventsConfiguration"`
	// Configuration for forwarding a partner event source's events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_source#partner_events_configuration Eventsv2EventSource#partner_events_configuration}
	PartnerEventsConfiguration *Eventsv2EventSourceConfigurationPartnerEventsConfiguration `field:"optional" json:"partnerEventsConfiguration" yaml:"partnerEventsConfiguration"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberFilterConfigurationFilters struct {
	// The event pattern, as a JSON string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#pattern Eventsv2Subscriber#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Which part of the event the pattern is evaluated against: DATA (the event payload), METADATA (event metadata), or SYSTEM_METADATA (service-generated metadata).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#scope Eventsv2Subscriber#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}


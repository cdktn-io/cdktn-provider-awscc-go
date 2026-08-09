// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventseventbus


type EventsEventBusLogConfig struct {
	// Configures whether or not to include event detail, input transformer details, target properties, and target input in the applicable log messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/events_event_bus#include_detail EventsEventBus#include_detail}
	IncludeDetail *string `field:"optional" json:"includeDetail" yaml:"includeDetail"`
	// Configures the log level of the EventBus and determines which log messages are sent to Ingestion Hub for delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/events_event_bus#level EventsEventBus#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotlogging


type IotLoggingEventConfigurations struct {
	// The type of event to log. These include event types like Connect, Publish, and Disconnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_logging#event_type IotLogging#event_type}
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// CloudWatch Log Group for event-based logging.
	//
	// Specifies where log events should be sent. The log destination for event-based logging overrides default Log Group for the specified event type and applies to all resources associated with that event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_logging#log_destination IotLogging#log_destination}
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// The logging level for the specified event type. Determines the verbosity of log messages generated for this event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_logging#log_level IotLogging#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventsource


type Eventsv2EventSourceConfigurationAwsServiceEventsConfiguration struct {
	// A single AWS service source identifier, for example aws.s3. Wildcards and lists are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_source#aws_service Eventsv2EventSource#aws_service}
	AwsService *string `field:"optional" json:"awsService" yaml:"awsService"`
	// The destination for events that could not be forwarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_source#on_failure_configuration Eventsv2EventSource#on_failure_configuration}
	OnFailureConfiguration *Eventsv2EventSourceConfigurationAwsServiceEventsConfigurationOnFailureConfiguration `field:"optional" json:"onFailureConfiguration" yaml:"onFailureConfiguration"`
	// A filter pattern, as a JSON string, that defines which events from the specified AWS service are forwarded to the event bus.
	//
	// Do not include source, account, or region as top-level fields. If you do not specify a pattern, all events from the service are forwarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_source#pattern Eventsv2EventSource#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}


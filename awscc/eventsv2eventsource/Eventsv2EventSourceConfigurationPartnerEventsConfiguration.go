// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventsource


type Eventsv2EventSourceConfigurationPartnerEventsConfiguration struct {
	// The destination for events that could not be forwarded, covering both the forwarding target and the managed partner event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#on_failure_configuration Eventsv2EventSource#on_failure_configuration}
	OnFailureConfiguration *Eventsv2EventSourceConfigurationPartnerEventsConfigurationOnFailureConfiguration `field:"optional" json:"onFailureConfiguration" yaml:"onFailureConfiguration"`
	// The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt events on the managed partner event bus.
	//
	// The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN. If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt events on the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#partner_bus_kms_key_identifier Eventsv2EventSource#partner_bus_kms_key_identifier}
	PartnerBusKmsKeyIdentifier *string `field:"optional" json:"partnerBusKmsKeyIdentifier" yaml:"partnerBusKmsKeyIdentifier"`
	// The ARN of the partner event source to forward.
	//
	// The partner owns the event source, so the ARN's account segment is empty. Changing this property replaces the event source. Because Name and EventBusArn together identify an event source, and the replacement is created before the old resource is deleted, change Name in the same update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#partner_event_source_arn Eventsv2EventSource#partner_event_source_arn}
	PartnerEventSourceArn *string `field:"optional" json:"partnerEventSourceArn" yaml:"partnerEventSourceArn"`
	// A filter pattern, as a JSON string, that defines which events from the specified partner event source are forwarded to the event bus.
	//
	// If you do not specify a pattern, all events from the partner event source are forwarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#pattern Eventsv2EventSource#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}


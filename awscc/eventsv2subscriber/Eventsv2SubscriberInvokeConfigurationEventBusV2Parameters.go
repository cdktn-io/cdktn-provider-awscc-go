// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationEventBusV2Parameters struct {
	// Deduplication settings applied to the forwarded events on the downstream event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#deduplication_configuration Eventsv2Subscriber#deduplication_configuration}
	DeduplicationConfiguration *Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersDeduplicationConfiguration `field:"optional" json:"deduplicationConfiguration" yaml:"deduplicationConfiguration"`
	// Metadata forwarded with each event, as key-value string pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#metadata Eventsv2Subscriber#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// System metadata attached to each forwarded event, controlling FIFO ordering and deduplication on the downstream event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#system_metadata Eventsv2Subscriber#system_metadata}
	SystemMetadata *Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadata `field:"optional" json:"systemMetadata" yaml:"systemMetadata"`
}


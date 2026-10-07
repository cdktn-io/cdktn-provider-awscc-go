// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersDeduplicationConfiguration struct {
	// How duplicate events are detected: CONTENT_BASED deduplicates by a hash of the event content.
	//
	// To deduplicate by a caller-supplied token instead, omit DeduplicationConfiguration and set SystemMetadata.DeduplicationId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#deduplication_type Eventsv2Subscriber#deduplication_type}
	DeduplicationType *string `field:"optional" json:"deduplicationType" yaml:"deduplicationType"`
}


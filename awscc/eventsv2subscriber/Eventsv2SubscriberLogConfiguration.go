// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberLogConfiguration struct {
	// Whether the event payload is included in emitted log records: FULL includes it in every emitted record, and ON_ERROR_ONLY includes it only in error records.
	//
	// The default is ON_ERROR_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#include_payload Eventsv2Subscriber#include_payload}
	IncludePayload *string `field:"optional" json:"includePayload" yaml:"includePayload"`
	// The minimum log level: OFF (no logging), ERROR, or INFO.
	//
	// Records below this level are not emitted. The default is OFF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#level Eventsv2Subscriber#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
}


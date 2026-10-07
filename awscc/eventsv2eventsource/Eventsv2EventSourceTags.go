// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventsource


type Eventsv2EventSourceTags struct {
	// The tag key. Unique per resource; keys are case sensitive. No leading or trailing whitespace (interior whitespace is allowed).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#key Eventsv2EventSource#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value. May be empty. No leading or trailing whitespace (interior whitespace is allowed).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#value Eventsv2EventSource#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


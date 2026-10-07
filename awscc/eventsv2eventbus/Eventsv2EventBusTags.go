// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventbus


type Eventsv2EventBusTags struct {
	// The tag key.
	//
	// For each resource, each tag key must be unique and each key can have only one value; keys are case sensitive. A key cannot begin or end with a whitespace character; whitespace inside the key is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_bus#key Eventsv2EventBus#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// May be empty. A value cannot begin or end with a whitespace character; whitespace inside the value is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_bus#value Eventsv2EventBus#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


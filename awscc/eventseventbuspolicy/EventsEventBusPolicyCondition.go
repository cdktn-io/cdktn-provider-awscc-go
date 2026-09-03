// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventseventbuspolicy


type EventsEventBusPolicyCondition struct {
	// Specifies the value for the key. Currently, this must be the ID of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#key EventsEventBusPolicy#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Specifies the type of condition. Currently the only supported value is StringEquals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#type EventsEventBusPolicy#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the key for the condition. Currently the only supported key is aws:PrincipalOrgID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#value EventsEventBusPolicy#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


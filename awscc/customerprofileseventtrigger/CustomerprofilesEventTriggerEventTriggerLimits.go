// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofileseventtrigger


type CustomerprofilesEventTriggerEventTriggerLimits struct {
	// Specifies that an event will only trigger the destination if it is processed within a certain latency period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_event_trigger#event_expiration CustomerprofilesEventTrigger#event_expiration}
	EventExpiration *float64 `field:"optional" json:"eventExpiration" yaml:"eventExpiration"`
	// A list of time periods during which the limits apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_event_trigger#periods CustomerprofilesEventTrigger#periods}
	Periods interface{} `field:"optional" json:"periods" yaml:"periods"`
}


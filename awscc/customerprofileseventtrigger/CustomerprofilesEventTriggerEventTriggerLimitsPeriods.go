// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofileseventtrigger


type CustomerprofilesEventTriggerEventTriggerLimitsPeriods struct {
	// The maximum allowed number of destination invocations per profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_event_trigger#max_invocations_per_profile CustomerprofilesEventTrigger#max_invocations_per_profile}
	MaxInvocationsPerProfile *float64 `field:"optional" json:"maxInvocationsPerProfile" yaml:"maxInvocationsPerProfile"`
	// The unit of time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_event_trigger#unit CustomerprofilesEventTrigger#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// If set to true, there is no limit on the number of destination invocations per profile.
	//
	// The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_event_trigger#unlimited CustomerprofilesEventTrigger#unlimited}
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
	// The amount of time of the specified unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_event_trigger#value CustomerprofilesEventTrigger#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


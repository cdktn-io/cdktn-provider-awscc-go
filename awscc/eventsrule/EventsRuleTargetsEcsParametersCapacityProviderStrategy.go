// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule


type EventsRuleTargetsEcsParametersCapacityProviderStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/events_rule#base EventsRule#base}.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/events_rule#capacity_provider EventsRule#capacity_provider}.
	CapacityProvider *string `field:"optional" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/events_rule#weight EventsRule#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


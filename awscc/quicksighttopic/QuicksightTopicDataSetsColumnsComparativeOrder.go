// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsColumnsComparativeOrder struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#specifed_order QuicksightTopic#specifed_order}.
	SpecifedOrder *[]*string `field:"optional" json:"specifedOrder" yaml:"specifedOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#treat_undefined_specified_values QuicksightTopic#treat_undefined_specified_values}.
	TreatUndefinedSpecifiedValues *string `field:"optional" json:"treatUndefinedSpecifiedValues" yaml:"treatUndefinedSpecifiedValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#use_ordering QuicksightTopic#use_ordering}.
	UseOrdering *string `field:"optional" json:"useOrdering" yaml:"useOrdering"`
}


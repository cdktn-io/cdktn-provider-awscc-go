// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsFiltersDateRangeFilterConstant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#constant_type QuicksightTopic#constant_type}.
	ConstantType *string `field:"optional" json:"constantType" yaml:"constantType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#range_constant QuicksightTopic#range_constant}.
	RangeConstant *QuicksightTopicDataSetsFiltersDateRangeFilterConstantRangeConstant `field:"optional" json:"rangeConstant" yaml:"rangeConstant"`
}


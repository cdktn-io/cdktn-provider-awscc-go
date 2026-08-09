// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsFiltersDateRangeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_topic#constant QuicksightTopic#constant}.
	Constant *QuicksightTopicDataSetsFiltersDateRangeFilterConstant `field:"optional" json:"constant" yaml:"constant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_topic#inclusive QuicksightTopic#inclusive}.
	Inclusive interface{} `field:"optional" json:"inclusive" yaml:"inclusive"`
}


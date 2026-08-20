// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsFiltersCategoryFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#category_filter_function QuicksightTopic#category_filter_function}.
	CategoryFilterFunction *string `field:"optional" json:"categoryFilterFunction" yaml:"categoryFilterFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#category_filter_type QuicksightTopic#category_filter_type}.
	CategoryFilterType *string `field:"optional" json:"categoryFilterType" yaml:"categoryFilterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#constant QuicksightTopic#constant}.
	Constant *QuicksightTopicDataSetsFiltersCategoryFilterConstant `field:"optional" json:"constant" yaml:"constant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#inverse QuicksightTopic#inverse}.
	Inverse interface{} `field:"optional" json:"inverse" yaml:"inverse"`
}


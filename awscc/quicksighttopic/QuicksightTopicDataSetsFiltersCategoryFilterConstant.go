// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsFiltersCategoryFilterConstant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_topic#collective_constant QuicksightTopic#collective_constant}.
	CollectiveConstant *QuicksightTopicDataSetsFiltersCategoryFilterConstantCollectiveConstant `field:"optional" json:"collectiveConstant" yaml:"collectiveConstant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_topic#constant_type QuicksightTopic#constant_type}.
	ConstantType *string `field:"optional" json:"constantType" yaml:"constantType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_topic#singular_constant QuicksightTopic#singular_constant}.
	SingularConstant *string `field:"optional" json:"singularConstant" yaml:"singularConstant"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsColumnsDefaultFormatting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_topic#display_format QuicksightTopic#display_format}.
	DisplayFormat *string `field:"optional" json:"displayFormat" yaml:"displayFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_topic#display_format_options QuicksightTopic#display_format_options}.
	DisplayFormatOptions *QuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptions `field:"optional" json:"displayFormatOptions" yaml:"displayFormatOptions"`
}


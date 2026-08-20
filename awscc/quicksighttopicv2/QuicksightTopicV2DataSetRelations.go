// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopicv2


type QuicksightTopicV2DataSetRelations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic_v2#left QuicksightTopicV2#left}.
	Left *QuicksightTopicV2DataSetRelationsLeft `field:"optional" json:"left" yaml:"left"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic_v2#right QuicksightTopicV2#right}.
	Right *QuicksightTopicV2DataSetRelationsRight `field:"optional" json:"right" yaml:"right"`
}


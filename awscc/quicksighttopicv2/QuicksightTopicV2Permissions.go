// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopicv2


type QuicksightTopicV2Permissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_topic_v2#actions QuicksightTopicV2#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_topic_v2#principal QuicksightTopicV2#principal}.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}


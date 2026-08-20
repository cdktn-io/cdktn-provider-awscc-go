// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopicv2


type QuicksightTopicV2DataSetRelationsLeft struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic_v2#column_names QuicksightTopicV2#column_names}.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic_v2#data_set_arn QuicksightTopicV2#data_set_arn}.
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
}


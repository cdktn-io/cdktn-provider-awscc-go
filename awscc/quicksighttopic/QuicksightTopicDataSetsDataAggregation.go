// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsDataAggregation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#dataset_row_date_granularity QuicksightTopic#dataset_row_date_granularity}.
	DatasetRowDateGranularity *string `field:"optional" json:"datasetRowDateGranularity" yaml:"datasetRowDateGranularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#default_date_column_name QuicksightTopic#default_date_column_name}.
	DefaultDateColumnName *string `field:"optional" json:"defaultDateColumnName" yaml:"defaultDateColumnName"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#calculated_fields QuicksightTopic#calculated_fields}.
	CalculatedFields interface{} `field:"optional" json:"calculatedFields" yaml:"calculatedFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#columns QuicksightTopic#columns}.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#data_aggregation QuicksightTopic#data_aggregation}.
	DataAggregation *QuicksightTopicDataSetsDataAggregation `field:"optional" json:"dataAggregation" yaml:"dataAggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#dataset_arn QuicksightTopic#dataset_arn}.
	DatasetArn *string `field:"optional" json:"datasetArn" yaml:"datasetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#dataset_description QuicksightTopic#dataset_description}.
	DatasetDescription *string `field:"optional" json:"datasetDescription" yaml:"datasetDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#dataset_name QuicksightTopic#dataset_name}.
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#filters QuicksightTopic#filters}.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#named_entities QuicksightTopic#named_entities}.
	NamedEntities interface{} `field:"optional" json:"namedEntities" yaml:"namedEntities"`
}


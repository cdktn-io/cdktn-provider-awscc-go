// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsNamedEntitiesDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#field_name QuicksightTopic#field_name}.
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#metric QuicksightTopic#metric}.
	Metric *QuicksightTopicDataSetsNamedEntitiesDefinitionMetric `field:"optional" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#property_name QuicksightTopic#property_name}.
	PropertyName *string `field:"optional" json:"propertyName" yaml:"propertyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#property_role QuicksightTopic#property_role}.
	PropertyRole *string `field:"optional" json:"propertyRole" yaml:"propertyRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#property_usage QuicksightTopic#property_usage}.
	PropertyUsage *string `field:"optional" json:"propertyUsage" yaml:"propertyUsage"`
}


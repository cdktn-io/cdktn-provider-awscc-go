// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsColumnsSemanticType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#falsey_cell_value QuicksightTopic#falsey_cell_value}.
	FalseyCellValue *string `field:"optional" json:"falseyCellValue" yaml:"falseyCellValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#falsey_cell_value_synonyms QuicksightTopic#falsey_cell_value_synonyms}.
	FalseyCellValueSynonyms *[]*string `field:"optional" json:"falseyCellValueSynonyms" yaml:"falseyCellValueSynonyms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#sub_type_name QuicksightTopic#sub_type_name}.
	SubTypeName *string `field:"optional" json:"subTypeName" yaml:"subTypeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#truthy_cell_value QuicksightTopic#truthy_cell_value}.
	TruthyCellValue *string `field:"optional" json:"truthyCellValue" yaml:"truthyCellValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#truthy_cell_value_synonyms QuicksightTopic#truthy_cell_value_synonyms}.
	TruthyCellValueSynonyms *[]*string `field:"optional" json:"truthyCellValueSynonyms" yaml:"truthyCellValueSynonyms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#type_name QuicksightTopic#type_name}.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#type_parameters QuicksightTopic#type_parameters}.
	TypeParameters *map[string]*string `field:"optional" json:"typeParameters" yaml:"typeParameters"`
}


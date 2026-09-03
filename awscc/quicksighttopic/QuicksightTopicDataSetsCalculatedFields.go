// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsCalculatedFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#aggregation QuicksightTopic#aggregation}.
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#allowed_aggregations QuicksightTopic#allowed_aggregations}.
	AllowedAggregations *[]*string `field:"optional" json:"allowedAggregations" yaml:"allowedAggregations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#calculated_field_description QuicksightTopic#calculated_field_description}.
	CalculatedFieldDescription *string `field:"optional" json:"calculatedFieldDescription" yaml:"calculatedFieldDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#calculated_field_name QuicksightTopic#calculated_field_name}.
	CalculatedFieldName *string `field:"optional" json:"calculatedFieldName" yaml:"calculatedFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#calculated_field_synonyms QuicksightTopic#calculated_field_synonyms}.
	CalculatedFieldSynonyms *[]*string `field:"optional" json:"calculatedFieldSynonyms" yaml:"calculatedFieldSynonyms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#cell_value_synonyms QuicksightTopic#cell_value_synonyms}.
	CellValueSynonyms interface{} `field:"optional" json:"cellValueSynonyms" yaml:"cellValueSynonyms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#column_data_role QuicksightTopic#column_data_role}.
	ColumnDataRole *string `field:"optional" json:"columnDataRole" yaml:"columnDataRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#comparative_order QuicksightTopic#comparative_order}.
	ComparativeOrder *QuicksightTopicDataSetsCalculatedFieldsComparativeOrder `field:"optional" json:"comparativeOrder" yaml:"comparativeOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#default_formatting QuicksightTopic#default_formatting}.
	DefaultFormatting *QuicksightTopicDataSetsCalculatedFieldsDefaultFormatting `field:"optional" json:"defaultFormatting" yaml:"defaultFormatting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#disable_indexing QuicksightTopic#disable_indexing}.
	DisableIndexing interface{} `field:"optional" json:"disableIndexing" yaml:"disableIndexing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#expression QuicksightTopic#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#is_included_in_topic QuicksightTopic#is_included_in_topic}.
	IsIncludedInTopic interface{} `field:"optional" json:"isIncludedInTopic" yaml:"isIncludedInTopic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#never_aggregate_in_filter QuicksightTopic#never_aggregate_in_filter}.
	NeverAggregateInFilter interface{} `field:"optional" json:"neverAggregateInFilter" yaml:"neverAggregateInFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#non_additive QuicksightTopic#non_additive}.
	NonAdditive interface{} `field:"optional" json:"nonAdditive" yaml:"nonAdditive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#not_allowed_aggregations QuicksightTopic#not_allowed_aggregations}.
	NotAllowedAggregations *[]*string `field:"optional" json:"notAllowedAggregations" yaml:"notAllowedAggregations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#semantic_type QuicksightTopic#semantic_type}.
	SemanticType *QuicksightTopicDataSetsCalculatedFieldsSemanticType `field:"optional" json:"semanticType" yaml:"semanticType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#time_granularity QuicksightTopic#time_granularity}.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}


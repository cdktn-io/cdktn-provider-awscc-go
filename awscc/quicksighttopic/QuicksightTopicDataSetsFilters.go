// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#category_filter QuicksightTopic#category_filter}.
	CategoryFilter *QuicksightTopicDataSetsFiltersCategoryFilter `field:"optional" json:"categoryFilter" yaml:"categoryFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#date_range_filter QuicksightTopic#date_range_filter}.
	DateRangeFilter *QuicksightTopicDataSetsFiltersDateRangeFilter `field:"optional" json:"dateRangeFilter" yaml:"dateRangeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#filter_class QuicksightTopic#filter_class}.
	FilterClass *string `field:"optional" json:"filterClass" yaml:"filterClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#filter_description QuicksightTopic#filter_description}.
	FilterDescription *string `field:"optional" json:"filterDescription" yaml:"filterDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#filter_name QuicksightTopic#filter_name}.
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#filter_synonyms QuicksightTopic#filter_synonyms}.
	FilterSynonyms *[]*string `field:"optional" json:"filterSynonyms" yaml:"filterSynonyms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#filter_type QuicksightTopic#filter_type}.
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#numeric_equality_filter QuicksightTopic#numeric_equality_filter}.
	NumericEqualityFilter *QuicksightTopicDataSetsFiltersNumericEqualityFilter `field:"optional" json:"numericEqualityFilter" yaml:"numericEqualityFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#numeric_range_filter QuicksightTopic#numeric_range_filter}.
	NumericRangeFilter *QuicksightTopicDataSetsFiltersNumericRangeFilter `field:"optional" json:"numericRangeFilter" yaml:"numericRangeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#operand_field_name QuicksightTopic#operand_field_name}.
	OperandFieldName *string `field:"optional" json:"operandFieldName" yaml:"operandFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#relative_date_filter QuicksightTopic#relative_date_filter}.
	RelativeDateFilter *QuicksightTopicDataSetsFiltersRelativeDateFilter `field:"optional" json:"relativeDateFilter" yaml:"relativeDateFilter"`
}


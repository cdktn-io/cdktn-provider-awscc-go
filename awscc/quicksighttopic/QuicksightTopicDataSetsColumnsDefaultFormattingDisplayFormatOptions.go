// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#blank_cell_format QuicksightTopic#blank_cell_format}.
	BlankCellFormat *string `field:"optional" json:"blankCellFormat" yaml:"blankCellFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#currency_symbol QuicksightTopic#currency_symbol}.
	CurrencySymbol *string `field:"optional" json:"currencySymbol" yaml:"currencySymbol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#date_format QuicksightTopic#date_format}.
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#decimal_separator QuicksightTopic#decimal_separator}.
	DecimalSeparator *string `field:"optional" json:"decimalSeparator" yaml:"decimalSeparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#fraction_digits QuicksightTopic#fraction_digits}.
	FractionDigits *float64 `field:"optional" json:"fractionDigits" yaml:"fractionDigits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#grouping_separator QuicksightTopic#grouping_separator}.
	GroupingSeparator *string `field:"optional" json:"groupingSeparator" yaml:"groupingSeparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#negative_format QuicksightTopic#negative_format}.
	NegativeFormat *QuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptionsNegativeFormat `field:"optional" json:"negativeFormat" yaml:"negativeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#prefix QuicksightTopic#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#suffix QuicksightTopic#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#unit_scaler QuicksightTopic#unit_scaler}.
	UnitScaler *string `field:"optional" json:"unitScaler" yaml:"unitScaler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#use_blank_cell_format QuicksightTopic#use_blank_cell_format}.
	UseBlankCellFormat interface{} `field:"optional" json:"useBlankCellFormat" yaml:"useBlankCellFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_topic#use_grouping QuicksightTopic#use_grouping}.
	UseGrouping interface{} `field:"optional" json:"useGrouping" yaml:"useGrouping"`
}


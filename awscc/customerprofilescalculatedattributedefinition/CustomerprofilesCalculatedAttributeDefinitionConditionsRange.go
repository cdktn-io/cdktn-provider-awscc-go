// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionConditionsRange struct {
	// The format the timestamp field in your JSON object is specified.
	//
	// This value should be one of EPOCHMILLI or ISO_8601. E.g. if your object type is MyType and source JSON is {"generatedAt": {"timestamp": "2001-07-04T12:08:56.235Z"}}, then TimestampFormat should be "ISO_8601".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_calculated_attribute_definition#timestamp_format CustomerprofilesCalculatedAttributeDefinition#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// An expression specifying the field in your JSON object from which the date should be parsed.
	//
	// The expression should follow the structure of \"{ObjectTypeName.<Location of timestamp field in JSON pointer format>}\". E.g. if your object type is MyType and source JSON is {"generatedAt": {"timestamp": "1737587945945"}}, then TimestampSource should be "{MyType.generatedAt.timestamp}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_calculated_attribute_definition#timestamp_source CustomerprofilesCalculatedAttributeDefinition#timestamp_source}
	TimestampSource *string `field:"optional" json:"timestampSource" yaml:"timestampSource"`
	// The unit of time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_calculated_attribute_definition#unit CustomerprofilesCalculatedAttributeDefinition#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The amount of time of the specified unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_calculated_attribute_definition#value CustomerprofilesCalculatedAttributeDefinition#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
	// A structure specifying the endpoints of the relative time period over which data is included in the aggregation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_calculated_attribute_definition#value_range CustomerprofilesCalculatedAttributeDefinition#value_range}
	ValueRange *CustomerprofilesCalculatedAttributeDefinitionConditionsRangeValueRange `field:"optional" json:"valueRange" yaml:"valueRange"`
}


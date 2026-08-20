// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisfirehose_delivery_stream#case_insensitive KinesisfirehoseDeliveryStream#case_insensitive}.
	CaseInsensitive interface{} `field:"optional" json:"caseInsensitive" yaml:"caseInsensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisfirehose_delivery_stream#column_to_json_key_mappings KinesisfirehoseDeliveryStream#column_to_json_key_mappings}.
	ColumnToJsonKeyMappings *map[string]*string `field:"optional" json:"columnToJsonKeyMappings" yaml:"columnToJsonKeyMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisfirehose_delivery_stream#convert_dots_in_json_keys_to_underscores KinesisfirehoseDeliveryStream#convert_dots_in_json_keys_to_underscores}.
	ConvertDotsInJsonKeysToUnderscores interface{} `field:"optional" json:"convertDotsInJsonKeysToUnderscores" yaml:"convertDotsInJsonKeysToUnderscores"`
}


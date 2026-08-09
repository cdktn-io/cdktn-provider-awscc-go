// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputSchemaRecordFormatMappingParametersCsvMappingParameters struct {
	// The column delimiter. For example, in a CSV format, a comma (",") is the typical column delimiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisanalyticsv2_application#record_column_delimiter Kinesisanalyticsv2Application#record_column_delimiter}
	RecordColumnDelimiter *string `field:"optional" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// The row delimiter. For example, in a CSV format, '\n' is the typical row delimiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisanalyticsv2_application#record_row_delimiter Kinesisanalyticsv2Application#record_row_delimiter}
	RecordRowDelimiter *string `field:"optional" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}


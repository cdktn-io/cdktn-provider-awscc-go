// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputSchema struct {
	// A list of `RecordColumn` objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#record_columns Kinesisanalyticsv2Application#record_columns}
	RecordColumns interface{} `field:"optional" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the encoding of the records in the streaming source. For example, UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#record_encoding Kinesisanalyticsv2Application#record_encoding}
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
	// Specifies the format of the records on the streaming source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#record_format Kinesisanalyticsv2Application#record_format}
	RecordFormat *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputSchemaRecordFormat `field:"optional" json:"recordFormat" yaml:"recordFormat"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputSchemaRecordColumns struct {
	// A reference to the data element in the streaming input or the reference data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#mapping Kinesisanalyticsv2Application#mapping}
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
	// The name of the column that is created in the in-application input stream or reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#name Kinesisanalyticsv2Application#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of column created in the in-application input stream or reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#sql_type Kinesisanalyticsv2Application#sql_type}
	SqlType *string `field:"optional" json:"sqlType" yaml:"sqlType"`
}


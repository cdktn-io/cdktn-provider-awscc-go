// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetContentDeliveryRulesDestinationS3DestinationConfigurationGlueConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#database_name IotanalyticsDataset#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#table_name IotanalyticsDataset#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}


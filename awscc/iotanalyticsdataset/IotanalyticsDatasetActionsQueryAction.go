// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetActionsQueryAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#filters IotanalyticsDataset#filters}.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#sql_query IotanalyticsDataset#sql_query}.
	SqlQuery *string `field:"optional" json:"sqlQuery" yaml:"sqlQuery"`
}


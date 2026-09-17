// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetActionsQueryActionFiltersDeltaTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_dataset#offset_seconds IotanalyticsDataset#offset_seconds}.
	OffsetSeconds *float64 `field:"optional" json:"offsetSeconds" yaml:"offsetSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_dataset#time_expression IotanalyticsDataset#time_expression}.
	TimeExpression *string `field:"optional" json:"timeExpression" yaml:"timeExpression"`
}


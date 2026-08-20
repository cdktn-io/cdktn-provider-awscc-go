// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetRetentionPeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#number_of_days IotanalyticsDataset#number_of_days}.
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#unlimited IotanalyticsDataset#unlimited}.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}


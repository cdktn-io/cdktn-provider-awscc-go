// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdatastore


type IotanalyticsDatastoreRetentionPeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_datastore#number_of_days IotanalyticsDatastore#number_of_days}.
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_datastore#unlimited IotanalyticsDatastore#unlimited}.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}


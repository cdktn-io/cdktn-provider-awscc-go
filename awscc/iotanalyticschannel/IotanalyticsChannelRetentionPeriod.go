// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticschannel


type IotanalyticsChannelRetentionPeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotanalytics_channel#number_of_days IotanalyticsChannel#number_of_days}.
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotanalytics_channel#unlimited IotanalyticsChannel#unlimited}.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}


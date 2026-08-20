// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobSchedulingConfigMaintenanceWindows struct {
	// Displays the duration of the next maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_job#duration_in_minutes IotJob#duration_in_minutes}
	DurationInMinutes *float64 `field:"optional" json:"durationInMinutes" yaml:"durationInMinutes"`
	// Displays the start time of the next maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_job#start_time IotJob#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}


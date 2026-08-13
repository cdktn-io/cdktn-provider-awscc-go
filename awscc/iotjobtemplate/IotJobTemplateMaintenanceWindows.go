// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjobtemplate


type IotJobTemplateMaintenanceWindows struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job_template#duration_in_minutes IotJobTemplate#duration_in_minutes}.
	DurationInMinutes *float64 `field:"optional" json:"durationInMinutes" yaml:"durationInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job_template#start_time IotJobTemplate#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}


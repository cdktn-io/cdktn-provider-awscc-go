// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobSchedulingConfig struct {
	// Specifies the end behavior for all job executions after a job reaches the selected endTime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job#end_behavior IotJob#end_behavior}
	EndBehavior *string `field:"optional" json:"endBehavior" yaml:"endBehavior"`
	// The time a job will stop rollout of the job document to all devices in the target group for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job#end_time IotJob#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// An optional configuration within the SchedulingConfig to setup a recurring maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job#maintenance_windows IotJob#maintenance_windows}
	MaintenanceWindows interface{} `field:"optional" json:"maintenanceWindows" yaml:"maintenanceWindows"`
	// The time a job will begin rollout of the job document to all devices in the target group for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job#start_time IotJob#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}


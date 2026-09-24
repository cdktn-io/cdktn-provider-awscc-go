// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetEcsParametersNetworkConfiguration struct {
	// This structure specifies the VPC subnets and security groups for the task, and whether a public IP address is to be used.
	//
	// This structure is relevant only for ECS tasks that use the awsvpc network mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#awsvpc_configuration SchedulerSchedule#awsvpc_configuration}
	AwsvpcConfiguration *SchedulerScheduleTargetEcsParametersNetworkConfigurationAwsvpcConfiguration `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
}


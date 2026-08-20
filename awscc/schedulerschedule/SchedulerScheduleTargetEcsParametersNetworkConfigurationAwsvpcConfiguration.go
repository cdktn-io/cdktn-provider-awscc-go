// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetEcsParametersNetworkConfigurationAwsvpcConfiguration struct {
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#assign_public_ip SchedulerSchedule#assign_public_ip}
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Specifies the security groups associated with the task.
	//
	// These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#security_groups SchedulerSchedule#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specifies the subnets associated with the task.
	//
	// These subnets must all be in the same VPC. You can specify as many as 16 subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#subnets SchedulerSchedule#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}


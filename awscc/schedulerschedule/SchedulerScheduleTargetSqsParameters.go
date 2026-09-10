// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetSqsParameters struct {
	// The FIFO message group ID to use as the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/scheduler_schedule#message_group_id SchedulerSchedule#message_group_id}
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}


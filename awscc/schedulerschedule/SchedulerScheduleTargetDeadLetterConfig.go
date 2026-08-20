// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetDeadLetterConfig struct {
	// The ARN of the SQS queue specified as the target for the dead-letter queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/scheduler_schedule#arn SchedulerSchedule#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}


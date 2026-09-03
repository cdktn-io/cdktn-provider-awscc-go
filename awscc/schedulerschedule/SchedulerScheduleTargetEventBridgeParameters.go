// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetEventBridgeParameters struct {
	// Free-form string, with a maximum of 128 characters, used to decide what fields to expect in the event detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/scheduler_schedule#detail_type SchedulerSchedule#detail_type}
	DetailType *string `field:"optional" json:"detailType" yaml:"detailType"`
	// The source of the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/scheduler_schedule#source SchedulerSchedule#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}


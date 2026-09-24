// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedulegroup


type SchedulerScheduleGroupTags struct {
	// Key for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule_group#key SchedulerScheduleGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule_group#value SchedulerScheduleGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


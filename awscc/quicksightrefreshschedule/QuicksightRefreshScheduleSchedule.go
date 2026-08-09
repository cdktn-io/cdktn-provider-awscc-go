// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightrefreshschedule


type QuicksightRefreshScheduleSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_refresh_schedule#refresh_type QuicksightRefreshSchedule#refresh_type}.
	RefreshType *string `field:"optional" json:"refreshType" yaml:"refreshType"`
	// <p>Information about the schedule frequency.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_refresh_schedule#schedule_frequency QuicksightRefreshSchedule#schedule_frequency}
	ScheduleFrequency *QuicksightRefreshScheduleScheduleScheduleFrequency `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// <p>An unique identifier for the refresh schedule.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_refresh_schedule#schedule_id QuicksightRefreshSchedule#schedule_id}
	ScheduleId *string `field:"optional" json:"scheduleId" yaml:"scheduleId"`
	// <p>The date time after which refresh is to be scheduled</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_refresh_schedule#start_after_date_time QuicksightRefreshSchedule#start_after_date_time}
	StartAfterDateTime *string `field:"optional" json:"startAfterDateTime" yaml:"startAfterDateTime"`
}


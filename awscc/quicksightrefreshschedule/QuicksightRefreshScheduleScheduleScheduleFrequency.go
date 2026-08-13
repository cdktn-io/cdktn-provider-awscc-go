// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightrefreshschedule


type QuicksightRefreshScheduleScheduleScheduleFrequency struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_refresh_schedule#interval QuicksightRefreshSchedule#interval}.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// <p>The day scheduled for refresh.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_refresh_schedule#refresh_on_day QuicksightRefreshSchedule#refresh_on_day}
	RefreshOnDay *QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDay `field:"optional" json:"refreshOnDay" yaml:"refreshOnDay"`
	// <p>The time of the day for scheduled refresh.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_refresh_schedule#time_of_the_day QuicksightRefreshSchedule#time_of_the_day}
	TimeOfTheDay *string `field:"optional" json:"timeOfTheDay" yaml:"timeOfTheDay"`
	// <p>The timezone for scheduled refresh.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_refresh_schedule#time_zone QuicksightRefreshSchedule#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}


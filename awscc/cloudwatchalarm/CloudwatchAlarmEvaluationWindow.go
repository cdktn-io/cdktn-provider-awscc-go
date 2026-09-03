// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmEvaluationWindow struct {
	// A sliding window, which advances each time the alarm is evaluated, forming a rolling time window.
	//
	// This is the default evaluation window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#sliding_window CloudwatchAlarm#sliding_window}
	SlidingWindow *string `field:"optional" json:"slidingWindow" yaml:"slidingWindow"`
	// A wall clock window, which aligns the evaluated range to fixed clock boundaries that match the alarm's period, such as the top of the hour, midnight, or the start of the calendar week.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#wall_clock_window CloudwatchAlarm#wall_clock_window}
	WallClockWindow *CloudwatchAlarmEvaluationWindowWallClockWindow `field:"optional" json:"wallClockWindow" yaml:"wallClockWindow"`
}


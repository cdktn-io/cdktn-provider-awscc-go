// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmEvaluationWindow struct {
	// Configuration for sliding evaluation window (default behavior).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudwatch_alarm#sliding_window CloudwatchAlarm#sliding_window}
	SlidingWindow *string `field:"optional" json:"slidingWindow" yaml:"slidingWindow"`
	// Configuration for wall clock based evaluation window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudwatch_alarm#wall_clock_window CloudwatchAlarm#wall_clock_window}
	WallClockWindow *CloudwatchAlarmEvaluationWindowWallClockWindow `field:"optional" json:"wallClockWindow" yaml:"wallClockWindow"`
}


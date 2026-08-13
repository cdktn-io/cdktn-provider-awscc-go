// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogalarm


type CloudwatchLogAlarmScheduledQueryConfigurationScheduleConfiguration struct {
	// The expression that defines when the scheduled query runs, e.g. rate(1 minute).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudwatch_log_alarm#schedule_expression CloudwatchLogAlarm#schedule_expression}
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The number of seconds into the past to start the query window.
	//
	// Must be a positive value and cannot exceed 2592000 seconds (30 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudwatch_log_alarm#start_time_offset CloudwatchLogAlarm#start_time_offset}
	StartTimeOffset *float64 `field:"required" json:"startTimeOffset" yaml:"startTimeOffset"`
	// The number of seconds into the past to end the query window.
	//
	// Must be a non-negative value and cannot exceed 2592000 seconds (30 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudwatch_log_alarm#end_time_offset CloudwatchLogAlarm#end_time_offset}
	EndTimeOffset *float64 `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
}


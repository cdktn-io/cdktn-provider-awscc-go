// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmWarmUpConfiguration struct {
	// Specifies whether the alarm waits for the full warm-up period before it starts to evaluate.
	//
	// The default is ``false``. If ``true``, the alarm waits the entire ``WarmUpPeriodDurationInMinutes`` before it starts to evaluate, even if metric data arrives earlier. If ``false``, the alarm ends the warm-up period early. Evaluation begins as soon as the alarm has enough metric data to fill its evaluation window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudwatch_alarm#only_start_evaluating_after_warm_up_period_ends CloudwatchAlarm#only_start_evaluating_after_warm_up_period_ends}
	OnlyStartEvaluatingAfterWarmUpPeriodEnds interface{} `field:"optional" json:"onlyStartEvaluatingAfterWarmUpPeriodEnds" yaml:"onlyStartEvaluatingAfterWarmUpPeriodEnds"`
	// The length of the warm-up period, in minutes.
	//
	// After you create or update the alarm, the alarm stays in ``INSUFFICIENT_DATA`` for this duration. During this time, the alarm does not perform alarm actions.
	//  You can change this value at any time, including after the warm-up period ends. If you change it after the warm-up period ends, the new value does not restart the warm-up period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudwatch_alarm#warm_up_period_duration_in_minutes CloudwatchAlarm#warm_up_period_duration_in_minutes}
	WarmUpPeriodDurationInMinutes *float64 `field:"optional" json:"warmUpPeriodDurationInMinutes" yaml:"warmUpPeriodDurationInMinutes"`
}


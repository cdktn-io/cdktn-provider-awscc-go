// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamscheduledquery


type TimestreamScheduledQueryScheduleConfiguration struct {
	// An expression that denotes when to trigger the scheduled query run.
	//
	// This can be a cron expression or a rate expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/timestream_scheduled_query#schedule_expression TimestreamScheduledQuery#schedule_expression}
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmscheduledreport


type BcmScheduledReportScheduleConfig struct {
	// The schedule expression that specifies when to trigger the scheduled report run.
	//
	// This value must be a cron expression consisting of six fields separated by white spaces: cron(minutes hours day_of_month month day_of_week year).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#schedule_expression BcmScheduledReport#schedule_expression}
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The time zone for the schedule expression, for example, UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#schedule_expression_time_zone BcmScheduledReport#schedule_expression_time_zone}
	ScheduleExpressionTimeZone *string `field:"optional" json:"scheduleExpressionTimeZone" yaml:"scheduleExpressionTimeZone"`
	// The time period during which the schedule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#schedule_period BcmScheduledReport#schedule_period}
	SchedulePeriod *BcmScheduledReportScheduleConfigSchedulePeriod `field:"optional" json:"schedulePeriod" yaml:"schedulePeriod"`
	// The state of the schedule.
	//
	// ENABLED means the scheduled report runs according to its schedule expression. DISABLED means the scheduled report is paused and will not run until re-enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#state BcmScheduledReport#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}


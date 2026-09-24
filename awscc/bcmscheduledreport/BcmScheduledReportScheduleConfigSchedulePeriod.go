// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmscheduledreport


type BcmScheduledReportScheduleConfigSchedulePeriod struct {
	// The time at which the schedule stops being active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#end_time BcmScheduledReport#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The time at which the schedule becomes active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#start_time BcmScheduledReport#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}


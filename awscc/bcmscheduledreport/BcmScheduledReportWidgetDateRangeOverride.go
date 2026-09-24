// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmscheduledreport


type BcmScheduledReportWidgetDateRangeOverride struct {
	// The end of the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#end_time BcmScheduledReport#end_time}
	EndTime *BcmScheduledReportWidgetDateRangeOverrideEndTime `field:"optional" json:"endTime" yaml:"endTime"`
	// The start of the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#start_time BcmScheduledReport#start_time}
	StartTime *BcmScheduledReportWidgetDateRangeOverrideStartTime `field:"optional" json:"startTime" yaml:"startTime"`
}


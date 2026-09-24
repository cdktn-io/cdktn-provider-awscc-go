// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmscheduledreport


type BcmScheduledReportWidgetDateRangeOverrideEndTime struct {
	// Whether Value is an absolute date or a duration relative to now.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#type BcmScheduledReport#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The date, or an ISO 8601 duration when Type is RELATIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#value BcmScheduledReport#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


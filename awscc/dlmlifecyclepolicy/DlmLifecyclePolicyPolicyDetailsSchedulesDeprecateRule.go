// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRule struct {
	// If the schedule has a count-based retention rule, this parameter specifies the number of oldest AMIs to deprecate.
	//
	// The count must be less than or equal to the schedule's retention count, and it can't be greater than 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#count DlmLifecyclePolicy#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// If the schedule has an age-based retention rule, this parameter specifies the period after which to deprecate AMIs created by the schedule.
	//
	// The period must be less than or equal to the schedule's retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time in which to measure the **Interval**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}


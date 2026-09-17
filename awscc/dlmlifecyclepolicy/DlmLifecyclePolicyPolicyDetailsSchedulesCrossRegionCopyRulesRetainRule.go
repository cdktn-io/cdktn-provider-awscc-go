// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesRetainRule struct {
	// The amount of time to retain a cross-Region snapshot or AMI copy.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time for time-based retention.
	//
	// For example, to retain a cross-Region copy for 3 months, specify `Interval=3` and `IntervalUnit=MONTHS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xraysamplingrule


type XraySamplingRuleSamplingRuleRecordSamplingRuleSamplingRateBoost struct {
	// Time window (in minutes) in which only one sampling rate boost can be triggered.
	//
	// After a boost occurs, no further boosts are allowed until the next window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/xray_sampling_rule#cooldown_window_minutes XraySamplingRule#cooldown_window_minutes}
	CooldownWindowMinutes *float64 `field:"optional" json:"cooldownWindowMinutes" yaml:"cooldownWindowMinutes"`
	// The maximum sampling rate X-Ray will apply when it detects anomalies.
	//
	// X-Ray determines the appropriate rate between your baseline and the maximum, depending on anomaly activity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/xray_sampling_rule#max_rate XraySamplingRule#max_rate}
	MaxRate *float64 `field:"optional" json:"maxRate" yaml:"maxRate"`
}


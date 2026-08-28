// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsworkspace


type ApsWorkspaceWorkspaceConfiguration struct {
	// An array of label set and associated limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_workspace#limits_per_label_sets ApsWorkspace#limits_per_label_sets}
	LimitsPerLabelSets interface{} `field:"optional" json:"limitsPerLabelSets" yaml:"limitsPerLabelSets"`
	// The time window in seconds for accepting out-of-order samples.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_workspace#out_of_order_time_window_in_seconds ApsWorkspace#out_of_order_time_window_in_seconds}
	OutOfOrderTimeWindowInSeconds *float64 `field:"optional" json:"outOfOrderTimeWindowInSeconds" yaml:"outOfOrderTimeWindowInSeconds"`
	// How many days that metrics are retained in the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_workspace#retention_period_in_days ApsWorkspace#retention_period_in_days}
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// Duration in seconds to offset rule evaluation queries into the past.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_workspace#rule_query_offset_in_seconds ApsWorkspace#rule_query_offset_in_seconds}
	RuleQueryOffsetInSeconds *float64 `field:"optional" json:"ruleQueryOffsetInSeconds" yaml:"ruleQueryOffsetInSeconds"`
}


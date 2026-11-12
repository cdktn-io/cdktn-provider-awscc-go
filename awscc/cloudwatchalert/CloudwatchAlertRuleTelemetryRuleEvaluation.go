// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertRuleTelemetryRuleEvaluation struct {
	// The time between evaluations, in seconds. Valid values: 30, 60, 120, 300, 600, 900, 1800, 3600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#interval_seconds CloudwatchAlert#interval_seconds}
	IntervalSeconds *float64 `field:"required" json:"intervalSeconds" yaml:"intervalSeconds"`
	// How long a threshold breach must last before the alert changes state, in seconds.
	//
	// Must be a multiple of IntervalSeconds. 0 changes state immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#pending_duration_seconds CloudwatchAlert#pending_duration_seconds}
	PendingDurationSeconds *float64 `field:"optional" json:"pendingDurationSeconds" yaml:"pendingDurationSeconds"`
	// How long the result must stay within its thresholds before the alert returns to OK, in seconds.
	//
	// Must be a multiple of IntervalSeconds. 0 returns to OK immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#recovery_duration_seconds CloudwatchAlert#recovery_duration_seconds}
	RecoveryDurationSeconds *float64 `field:"optional" json:"recoveryDurationSeconds" yaml:"recoveryDurationSeconds"`
}


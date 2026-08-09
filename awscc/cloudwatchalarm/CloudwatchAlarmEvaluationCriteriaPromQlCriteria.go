// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmEvaluationCriteriaPromQlCriteria struct {
	// The duration, in seconds, that a contributor must be continuously breaching before it transitions to the ``ALARM`` state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudwatch_alarm#pending_period CloudwatchAlarm#pending_period}
	PendingPeriod *float64 `field:"optional" json:"pendingPeriod" yaml:"pendingPeriod"`
	// The PromQL query that the alarm evaluates.
	//
	// The query must return a result of vector type. Each entry in the vector result represents an alarm contributor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudwatch_alarm#query CloudwatchAlarm#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The duration, in seconds, that a contributor must continuously not be breaching before it transitions back to the ``OK`` state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudwatch_alarm#recovery_period CloudwatchAlarm#recovery_period}
	RecoveryPeriod *float64 `field:"optional" json:"recoveryPeriod" yaml:"recoveryPeriod"`
}


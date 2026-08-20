// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmEvaluationCriteria struct {
	// The PromQL criteria for the alarm evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_alarm#prom_ql_criteria CloudwatchAlarm#prom_ql_criteria}
	PromQlCriteria *CloudwatchAlarmEvaluationCriteriaPromQlCriteria `field:"optional" json:"promQlCriteria" yaml:"promQlCriteria"`
}


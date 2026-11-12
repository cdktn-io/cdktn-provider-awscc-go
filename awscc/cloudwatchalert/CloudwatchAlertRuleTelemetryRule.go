// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertRuleTelemetryRule struct {
	// The thresholds that set the alert's state. Specify at least one of WarningThreshold and CriticalThreshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#condition CloudwatchAlert#condition}
	Condition *CloudwatchAlertRuleTelemetryRuleCondition `field:"required" json:"condition" yaml:"condition"`
	// How often the alert is evaluated, and how long a change must last before the state changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#evaluation CloudwatchAlert#evaluation}
	Evaluation *CloudwatchAlertRuleTelemetryRuleEvaluation `field:"required" json:"evaluation" yaml:"evaluation"`
	// The query the alert runs, and the language it is written in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#query CloudwatchAlert#query}
	Query *CloudwatchAlertRuleTelemetryRuleQuery `field:"required" json:"query" yaml:"query"`
	// How the alert behaves when an evaluation returns no data. If omitted, the alert reports NODATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#no_data CloudwatchAlert#no_data}
	NoData *CloudwatchAlertRuleTelemetryRuleNoData `field:"optional" json:"noData" yaml:"noData"`
}


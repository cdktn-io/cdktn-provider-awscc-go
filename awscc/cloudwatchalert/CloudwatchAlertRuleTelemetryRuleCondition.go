// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertRuleTelemetryRuleCondition struct {
	// The comparison operator.
	//
	// Valid values: GT, GTE, LT, LTE (greater than, greater than or equal to, less than, less than or equal to).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#comparator CloudwatchAlert#comparator}
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// The value at which the alert enters the CRITICAL state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#critical_threshold CloudwatchAlert#critical_threshold}
	CriticalThreshold *float64 `field:"optional" json:"criticalThreshold" yaml:"criticalThreshold"`
	// The result field to compare against the threshold. Required when ThresholdMode is FIELD_VALUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#threshold_field CloudwatchAlert#threshold_field}
	ThresholdField *string `field:"optional" json:"thresholdField" yaml:"thresholdField"`
	// How the threshold is applied.
	//
	// COUNT_OF_RESULTS compares the number of rows the query returns. FIELD_VALUE compares the value of ThresholdField in each result row; each row is tracked separately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#threshold_mode CloudwatchAlert#threshold_mode}
	ThresholdMode *string `field:"optional" json:"thresholdMode" yaml:"thresholdMode"`
	// The value at which the alert enters the WARNING state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#warning_threshold CloudwatchAlert#warning_threshold}
	WarningThreshold *float64 `field:"optional" json:"warningThreshold" yaml:"warningThreshold"`
}


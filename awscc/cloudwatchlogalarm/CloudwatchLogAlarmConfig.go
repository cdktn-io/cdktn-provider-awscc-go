// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogalarm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchLogAlarmConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The arithmetic operation to use when comparing the specified threshold and the query results.
	//
	// Valid values are GreaterThanOrEqualToThreshold, GreaterThanThreshold, LessThanThreshold, and LessThanOrEqualToThreshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#comparison_operator CloudwatchLogAlarm#comparison_operator}
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of query results that must be breaching to trigger the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#query_results_to_alarm CloudwatchLogAlarm#query_results_to_alarm}
	QueryResultsToAlarm *float64 `field:"required" json:"queryResultsToAlarm" yaml:"queryResultsToAlarm"`
	// The number of query results over which data is compared to the specified threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#query_results_to_evaluate CloudwatchLogAlarm#query_results_to_evaluate}
	QueryResultsToEvaluate *float64 `field:"required" json:"queryResultsToEvaluate" yaml:"queryResultsToEvaluate"`
	// The scheduled query configuration for the log alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#scheduled_query_configuration CloudwatchLogAlarm#scheduled_query_configuration}
	ScheduledQueryConfiguration *CloudwatchLogAlarmScheduledQueryConfiguration `field:"required" json:"scheduledQueryConfiguration" yaml:"scheduledQueryConfiguration"`
	// The value to compare against the results of the scheduled query evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#threshold CloudwatchLogAlarm#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The number of log lines to include in alarm notifications. Valid values are 0 to 50.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#action_log_line_count CloudwatchLogAlarm#action_log_line_count}
	ActionLogLineCount *float64 `field:"optional" json:"actionLogLineCount" yaml:"actionLogLineCount"`
	// The ARN of the IAM role that grants CloudWatch permissions to fetch log lines for alarm notifications.
	//
	// Required when ActionLogLineCount is greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#action_log_line_role_arn CloudwatchLogAlarm#action_log_line_role_arn}
	ActionLogLineRoleArn *string `field:"optional" json:"actionLogLineRoleArn" yaml:"actionLogLineRoleArn"`
	// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#actions_enabled CloudwatchLogAlarm#actions_enabled}
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#alarm_actions CloudwatchLogAlarm#alarm_actions}
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// The description of the log alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#alarm_description CloudwatchLogAlarm#alarm_description}
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// The name of the log alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#alarm_name CloudwatchLogAlarm#alarm_name}
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#insufficient_data_actions CloudwatchLogAlarm#insufficient_data_actions}
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// The actions to execute when this alarm transitions to the OK state from any other state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#ok_actions CloudwatchLogAlarm#ok_actions}
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// A list of key-value pairs to associate with the log alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#tags CloudwatchLogAlarm#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudwatch_log_alarm#treat_missing_data CloudwatchLogAlarm#treat_missing_data}
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}


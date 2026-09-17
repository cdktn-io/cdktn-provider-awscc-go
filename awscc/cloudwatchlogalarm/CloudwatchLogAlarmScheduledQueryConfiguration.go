// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogalarm


type CloudwatchLogAlarmScheduledQueryConfiguration struct {
	// The aggregation expression for the scheduled query, e.g. count(*) or avg(latency) by host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_log_alarm#aggregation_expression CloudwatchLogAlarm#aggregation_expression}
	AggregationExpression *string `field:"required" json:"aggregationExpression" yaml:"aggregationExpression"`
	// The query string to execute against the specified log groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_log_alarm#query_string CloudwatchLogAlarm#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// The schedule configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_log_alarm#schedule_configuration CloudwatchLogAlarm#schedule_configuration}
	ScheduleConfiguration *CloudwatchLogAlarmScheduledQueryConfigurationScheduleConfiguration `field:"required" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The ARN of the IAM role that grants permissions to execute the scheduled query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_log_alarm#scheduled_query_role_arn CloudwatchLogAlarm#scheduled_query_role_arn}
	ScheduledQueryRoleArn *string `field:"required" json:"scheduledQueryRoleArn" yaml:"scheduledQueryRoleArn"`
	// The log groups to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_log_alarm#log_group_identifiers CloudwatchLogAlarm#log_group_identifiers}
	LogGroupIdentifiers *[]*string `field:"optional" json:"logGroupIdentifiers" yaml:"logGroupIdentifiers"`
	// A list of key-value pairs to associate with the scheduled query that backs the log alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_log_alarm#tags CloudwatchLogAlarm#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


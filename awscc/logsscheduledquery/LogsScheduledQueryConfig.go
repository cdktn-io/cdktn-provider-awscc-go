// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsscheduledquery

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsScheduledQueryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#execution_role_arn LogsScheduledQuery#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#name LogsScheduledQuery#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#query_language LogsScheduledQuery#query_language}.
	QueryLanguage *string `field:"required" json:"queryLanguage" yaml:"queryLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#query_string LogsScheduledQuery#query_string}.
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#schedule_expression LogsScheduledQuery#schedule_expression}.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#description LogsScheduledQuery#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#destination_configuration LogsScheduledQuery#destination_configuration}.
	DestinationConfiguration *LogsScheduledQueryDestinationConfiguration `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#log_group_identifiers LogsScheduledQuery#log_group_identifiers}.
	LogGroupIdentifiers *[]*string `field:"optional" json:"logGroupIdentifiers" yaml:"logGroupIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#schedule_end_time LogsScheduledQuery#schedule_end_time}.
	ScheduleEndTime *float64 `field:"optional" json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#schedule_start_time LogsScheduledQuery#schedule_start_time}.
	ScheduleStartTime *float64 `field:"optional" json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#start_time_offset LogsScheduledQuery#start_time_offset}.
	StartTimeOffset *float64 `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#state LogsScheduledQuery#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#tags LogsScheduledQuery#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_scheduled_query#timezone LogsScheduledQuery#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}


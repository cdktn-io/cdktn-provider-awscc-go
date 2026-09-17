// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncgraphqlapi


type AppsyncGraphQlApiLogConfig struct {
	// The service role that AWS AppSync will assume to publish to Amazon CloudWatch Logs in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appsync_graph_ql_api#cloudwatch_logs_role_arn AppsyncGraphQlApi#cloudwatch_logs_role_arn}
	CloudwatchLogsRoleArn *string `field:"optional" json:"cloudwatchLogsRoleArn" yaml:"cloudwatchLogsRoleArn"`
	// Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appsync_graph_ql_api#exclude_verbose_content AppsyncGraphQlApi#exclude_verbose_content}
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// The field logging level. Values can be NONE, ERROR, INFO, DEBUG, or ALL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appsync_graph_ql_api#field_log_level AppsyncGraphQlApi#field_log_level}
	FieldLogLevel *string `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
}


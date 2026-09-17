// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamscheduledquery

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TimestreamScheduledQueryConfig struct {
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
	// Configuration for error reporting. Error reports will be generated when a problem is encountered when writing the query results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#error_report_configuration TimestreamScheduledQuery#error_report_configuration}
	ErrorReportConfiguration *TimestreamScheduledQueryErrorReportConfiguration `field:"required" json:"errorReportConfiguration" yaml:"errorReportConfiguration"`
	// Notification configuration for the scheduled query.
	//
	// A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#notification_configuration TimestreamScheduledQuery#notification_configuration}
	NotificationConfiguration *TimestreamScheduledQueryNotificationConfiguration `field:"required" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// The query string to run.
	//
	// Parameter names can be specified in the query string @ character followed by an identifier. The named Parameter
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Configuration for when the scheduled query is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#schedule_configuration TimestreamScheduledQuery#schedule_configuration}
	ScheduleConfiguration *TimestreamScheduledQueryScheduleConfiguration `field:"required" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The ARN for the IAM role that Timestream will assume when running the scheduled query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#scheduled_query_execution_role_arn TimestreamScheduledQuery#scheduled_query_execution_role_arn}
	ScheduledQueryExecutionRoleArn *string `field:"required" json:"scheduledQueryExecutionRoleArn" yaml:"scheduledQueryExecutionRoleArn"`
	// Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words, making the same request repeatedly will produce the same result.
	//
	// Making multiple identical CreateScheduledQuery requests has the same effect as making a single request. If CreateScheduledQuery is called without a ClientToken, the Query SDK generates a ClientToken on your behalf. After 8 hours, any request with the same ClientToken is treated as a new request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#client_token TimestreamScheduledQuery#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The Amazon KMS key used to encrypt the scheduled query resource, at-rest.
	//
	// If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with alias/. If ErrorReportConfiguration uses SSE_KMS as encryption type, the same KmsKeyId is used to encrypt the error report at rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#kms_key_id TimestreamScheduledQuery#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the scheduled query. Scheduled query names must be unique within each Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#scheduled_query_name TimestreamScheduledQuery#scheduled_query_name}
	ScheduledQueryName *string `field:"optional" json:"scheduledQueryName" yaml:"scheduledQueryName"`
	// A list of key-value pairs to label the scheduled query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#tags TimestreamScheduledQuery#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Configuration of target store where scheduled query results are written to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#target_configuration TimestreamScheduledQuery#target_configuration}
	TargetConfiguration *TimestreamScheduledQueryTargetConfiguration `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logssubscriptionfilter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsSubscriptionFilterConfig struct {
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
	// The Amazon Resource Name (ARN) of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#destination_arn LogsSubscriptionFilter#destination_arn}
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The filtering expressions that restrict what gets delivered to the destination AWS resource.
	//
	// For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#filter_pattern LogsSubscriptionFilter#filter_pattern}
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The log group to associate with the subscription filter.
	//
	// All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#log_group_name LogsSubscriptionFilter#log_group_name}
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// This parameter is valid only for log groups that have an active log transformer.
	//
	// For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html).
	//  If this value is ``true``, the subscription filter is applied on the transformed version of the log events instead of the original ingested log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#apply_on_transformed_logs LogsSubscriptionFilter#apply_on_transformed_logs}
	ApplyOnTransformedLogs interface{} `field:"optional" json:"applyOnTransformedLogs" yaml:"applyOnTransformedLogs"`
	// The method used to distribute log data to the destination, which can be either random or grouped by log stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#distribution LogsSubscriptionFilter#distribution}
	Distribution *string `field:"optional" json:"distribution" yaml:"distribution"`
	// The list of system fields that are included in the log events sent to the subscription destination.
	//
	// Returns the ``emitSystemFields`` value if it was specified when the subscription filter was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#emit_system_fields LogsSubscriptionFilter#emit_system_fields}
	EmitSystemFields *[]*string `field:"optional" json:"emitSystemFields" yaml:"emitSystemFields"`
	// The filter expression that specifies which log events are processed by this subscription filter based on system fields.
	//
	// Returns the ``fieldSelectionCriteria`` value if it was specified when the subscription filter was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#field_selection_criteria LogsSubscriptionFilter#field_selection_criteria}
	FieldSelectionCriteria *string `field:"optional" json:"fieldSelectionCriteria" yaml:"fieldSelectionCriteria"`
	// The name of the subscription filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#filter_name LogsSubscriptionFilter#filter_name}
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// The ARN of an IAM role that grants CWL permissions to deliver ingested log events to the destination stream.
	//
	// You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_subscription_filter#role_arn LogsSubscriptionFilter#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}


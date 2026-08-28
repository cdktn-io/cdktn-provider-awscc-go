// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DynamodbTableConfig struct {
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
	// Specifies the attributes that make up the primary key for the table.
	//
	// The attributes in the ``KeySchema`` property must also be defined in the ``AttributeDefinitions`` property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#key_schema DynamodbTable#key_schema}
	KeySchema *string `field:"required" json:"keySchema" yaml:"keySchema"`
	// A list of attributes that describe the key schema for the table and indexes.
	//
	// This property is required to create a DDB table.
	//  Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Replacement if you edit an existing AttributeDefinition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#attribute_definitions DynamodbTable#attribute_definitions}
	AttributeDefinitions interface{} `field:"optional" json:"attributeDefinitions" yaml:"attributeDefinitions"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	//
	// Valid values include:
	//   +  ``PAY_PER_REQUEST`` - We recommend using ``PAY_PER_REQUEST`` for most DynamoDB workloads. ``PAY_PER_REQUEST`` sets the billing mode to [On-demand capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/on-demand-capacity-mode.html).
	//   +  ``PROVISIONED`` - We recommend using ``PROVISIONED`` for steady workloads with predictable growth where capacity requirements can be reliably forecasted. ``PROVISIONED`` sets the billing mode to [Provisioned capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/provisioned-capacity-mode.html).
	//
	//  If not specified, the default is ``PROVISIONED``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#billing_mode DynamodbTable#billing_mode}
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// The settings used to specify whether to enable CloudWatch Contributor Insights for the table and define which events to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#contributor_insights_specification DynamodbTable#contributor_insights_specification}
	ContributorInsightsSpecification *DynamodbTableContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Determines if a table is protected from deletion.
	//
	// When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#deletion_protection_enabled DynamodbTable#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Global secondary indexes to be created on the table.
	//
	// You can create up to 20 global secondary indexes.
	//   If you update a table to include a new global secondary index, CFNlong initiates the index creation and then proceeds with the stack update. CFNlong doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is ``ACTIVE``. You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.
	//  If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index.
	//  Updates are not supported. The following are exceptions:
	//   +  If you update either the contributor insights specification or the provisioned throughput values of global secondary indexes, you can update the table without interruption.
	//   +  You can delete or add one global secondary index without interruption. If you do both in the same update (for example, by changing the index's logical ID), the update fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#global_secondary_indexes DynamodbTable#global_secondary_indexes}
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Specifies the properties of data being imported from the S3 bucket source to the" table.
	//
	// If you specify the ``ImportSourceSpecification`` property, and also specify either the ``StreamSpecification``, the ``TableClass`` property, the ``DeletionProtectionEnabled`` property, or the ``WarmThroughput`` property, the IAM entity creating/updating stack must have ``UpdateTable`` permission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#import_source_specification DynamodbTable#import_source_specification}
	ImportSourceSpecification *DynamodbTableImportSourceSpecification `field:"optional" json:"importSourceSpecification" yaml:"importSourceSpecification"`
	// The Kinesis Data Streams configuration for the specified table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#kinesis_stream_specification DynamodbTable#kinesis_stream_specification}
	KinesisStreamSpecification *DynamodbTableKinesisStreamSpecification `field:"optional" json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// Local secondary indexes to be created on the table.
	//
	// You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#local_secondary_indexes DynamodbTable#local_secondary_indexes}
	LocalSecondaryIndexes interface{} `field:"optional" json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// Sets the maximum number of read and write units for the specified on-demand table.
	//
	// If you use this property, you must specify ``MaxReadRequestUnits``, ``MaxWriteRequestUnits``, or both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#on_demand_throughput DynamodbTable#on_demand_throughput}
	OnDemandThroughput *DynamodbTableOnDemandThroughput `field:"optional" json:"onDemandThroughput" yaml:"onDemandThroughput"`
	// The settings used to enable point in time recovery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#point_in_time_recovery_specification DynamodbTable#point_in_time_recovery_specification}
	PointInTimeRecoverySpecification *DynamodbTablePointInTimeRecoverySpecification `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``.
	//
	// For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html).
	//  If you set ``BillingMode`` as ``PROVISIONED``, you must specify this property. If you set ``BillingMode`` as ``PAY_PER_REQUEST``, you cannot specify this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#provisioned_throughput DynamodbTable#provisioned_throughput}
	ProvisionedThroughput *DynamodbTableProvisionedThroughput `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// An AWS resource-based policy document in JSON format that will be attached to the table.
	//
	// When you attach a resource-based policy while creating a table, the policy application is *strongly consistent*.
	//  The maximum size supported for a resource-based policy document is 20 KB. DynamoDB counts whitespaces when calculating the size of a policy against this limit. For a full list of all considerations that apply for resource-based policies, see [Resource-based policy considerations](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-considerations.html).
	//   You need to specify the ``CreateTable`` and ``PutResourcePolicy`` IAM actions for authorizing a user to create a table with a resource-based policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#resource_policy DynamodbTable#resource_policy}
	ResourcePolicy *DynamodbTableResourcePolicy `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// Specifies the settings to enable server-side encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#sse_specification DynamodbTable#sse_specification}
	SseSpecification *DynamodbTableSseSpecification `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The settings for the DDB table stream, which captures changes to items stored in the table.
	//
	// Including this property in your CFNlong template automatically enables streaming.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#stream_specification DynamodbTable#stream_specification}
	StreamSpecification *DynamodbTableStreamSpecification `field:"optional" json:"streamSpecification" yaml:"streamSpecification"`
	// The table class of the new table. Valid values are ``STANDARD`` and ``STANDARD_INFREQUENT_ACCESS``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#table_class DynamodbTable#table_class}
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// A name for the table.
	//
	// If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#table_name DynamodbTable#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// An array of key-value pairs to apply to this resource.  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#tags DynamodbTable#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the Time to Live (TTL) settings for the table.
	//
	// For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#time_to_live_specification DynamodbTable#time_to_live_specification}
	TimeToLiveSpecification *DynamodbTableTimeToLiveSpecification `field:"optional" json:"timeToLiveSpecification" yaml:"timeToLiveSpecification"`
	// Represents the warm throughput (in read units per second and write units per second) for creating a table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#warm_throughput DynamodbTable#warm_throughput}
	WarmThroughput *DynamodbTableWarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
}


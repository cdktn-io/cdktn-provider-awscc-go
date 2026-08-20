// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaEventSourceMappingConfig struct {
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
	// The name or ARN of the Lambda function.
	//
	// **Name formats**
	//  +  *Function name* – ``MyFunction``.
	//   +  *Function ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:MyFunction``.
	//   +  *Version or Alias ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD``.
	//   +  *Partial ARN* – ``123456789012:function:MyFunction``.
	//
	//  The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#function_name LambdaEventSourceMapping#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#amazon_managed_kafka_event_source_config LambdaEventSourceMapping#amazon_managed_kafka_event_source_config}
	AmazonManagedKafkaEventSourceConfig *LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfig `field:"optional" json:"amazonManagedKafkaEventSourceConfig" yaml:"amazonManagedKafkaEventSourceConfig"`
	// The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function.
	//
	// Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).
	//   +  *Amazon Kinesis* – Default 100. Max 10,000.
	//   +  *Amazon DynamoDB Streams* – Default 100. Max 10,000.
	//   +  *Amazon Simple Queue Service* – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.
	//   +  *Amazon Managed Streaming for Apache Kafka* – Default 100. Max 10,000.
	//   +  *Self-managed Apache Kafka* – Default 100. Max 10,000.
	//   +  *Amazon MQ (ActiveMQ and RabbitMQ)* – Default 100. Max 10,000.
	//   +  *DocumentDB* – Default 100. Max 10,000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#batch_size LambdaEventSourceMapping#batch_size}
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry.
	//
	// The default value is false.
	//   When using ``BisectBatchOnFunctionError``, check the ``BatchSize`` parameter in the ``OnFailure`` destination message's metadata. The ``BatchSize`` could be greater than 1 since LAM consolidates failed messages metadata when writing to the ``OnFailure`` destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#bisect_batch_on_function_error LambdaEventSourceMapping#bisect_batch_on_function_error}
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka) A configuration object that specifies the destination of an event after Lambda processes it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#destination_config LambdaEventSourceMapping#destination_config}
	DestinationConfig *LambdaEventSourceMappingDestinationConfig `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// Specific configuration settings for a DocumentDB event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#document_db_event_source_config LambdaEventSourceMapping#document_db_event_source_config}
	DocumentDbEventSourceConfig *LambdaEventSourceMappingDocumentDbEventSourceConfig `field:"optional" json:"documentDbEventSourceConfig" yaml:"documentDbEventSourceConfig"`
	// When true, the event source mapping is active. When false, Lambda pauses polling and invocation.  Default: True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#enabled LambdaEventSourceMapping#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	//
	// +  *Amazon Kinesis* – The ARN of the data stream or a stream consumer.
	//   +  *Amazon DynamoDB Streams* – The ARN of the stream.
	//   +  *Amazon Simple Queue Service* – The ARN of the queue.
	//   +  *Amazon Managed Streaming for Apache Kafka* – The ARN of the cluster or the ARN of the VPC connection (for [cross-account event source mappings](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#msk-multi-vpc)).
	//   +  *Amazon MQ* – The ARN of the broker.
	//   +  *Amazon DocumentDB* – The ARN of the DocumentDB change stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#event_source_arn LambdaEventSourceMapping#event_source_arn}
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// An object that defines the filter criteria that determine whether Lambda should process an event.
	//
	// For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#filter_criteria LambdaEventSourceMapping#filter_criteria}
	FilterCriteria *LambdaEventSourceMappingFilterCriteria `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// (Kinesis, DynamoDB Streams, and SQS) A list of current response type enums applied to the event source mapping.
	//
	// Valid Values: ``ReportBatchItemFailures``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#function_response_types LambdaEventSourceMapping#function_response_types}
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// The ARN of the KMSlong (KMS) customer managed key that Lambda uses to encrypt your function's [filter criteria](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#kms_key_arn LambdaEventSourceMapping#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// (Amazon MSK, and self-managed Apache Kafka only) The logging configuration for your event source.
	//
	// For more information, see [Event source mapping logging](https://docs.aws.amazon.com/lambda/latest/dg/esm-logging.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#logging_config LambdaEventSourceMapping#logging_config}
	LoggingConfig *LambdaEventSourceMappingLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.
	//
	// *Default (, , event sources)*: 0
	//  *Default (, Kafka, , event sources)*: 500 ms
	//  *Related setting:* For SQS event sources, when you set ``BatchSize`` to a value greater than 10, you must set ``MaximumBatchingWindowInSeconds`` to at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#maximum_batching_window_in_seconds LambdaEventSourceMapping#maximum_batching_window_in_seconds}
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka) Discard records older than the specified age.
	//
	// The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.
	//   The minimum valid value for maximum record age is 60s. Although values less than 60 and greater than -1 fall within the parameter's absolute range, they are not allowed
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#maximum_record_age_in_seconds LambdaEventSourceMapping#maximum_record_age_in_seconds}
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka) Discard records after the specified number of retries.
	//
	// The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#maximum_retry_attempts LambdaEventSourceMapping#maximum_retry_attempts}
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// The metrics configuration for your event source. For more information, see [Event source mapping metrics](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics-types.html#event-source-mapping-metrics).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#metrics_config LambdaEventSourceMapping#metrics_config}
	MetricsConfig *LambdaEventSourceMappingMetricsConfig `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// (Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard.
	//
	// The default value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#parallelization_factor LambdaEventSourceMapping#parallelization_factor}
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// (Amazon SQS, Amazon MSK, and self-managed Apache Kafka only) The provisioned mode configuration for the event source.
	//
	// For more information, see [provisioned mode](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-provisioned-mode).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#provisioned_poller_config LambdaEventSourceMapping#provisioned_poller_config}
	ProvisionedPollerConfig *LambdaEventSourceMappingProvisionedPollerConfig `field:"optional" json:"provisionedPollerConfig" yaml:"provisionedPollerConfig"`
	// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#queues LambdaEventSourceMapping#queues}
	Queues *[]*string `field:"optional" json:"queues" yaml:"queues"`
	// This property is for Amazon SQS event sources only.
	//
	// You cannot use ``ProvisionedPollerConfig`` while using ``ScalingConfig``. These options are mutually exclusive. To remove the scaling configuration, pass an empty value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#scaling_config LambdaEventSourceMapping#scaling_config}
	ScalingConfig *LambdaEventSourceMappingScalingConfig `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The self-managed Apache Kafka cluster for your event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#self_managed_event_source LambdaEventSourceMapping#self_managed_event_source}
	SelfManagedEventSource *LambdaEventSourceMappingSelfManagedEventSource `field:"optional" json:"selfManagedEventSource" yaml:"selfManagedEventSource"`
	// Specific configuration settings for a self-managed Apache Kafka event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#self_managed_kafka_event_source_config LambdaEventSourceMapping#self_managed_kafka_event_source_config}
	SelfManagedKafkaEventSourceConfig *LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig `field:"optional" json:"selfManagedKafkaEventSourceConfig" yaml:"selfManagedKafkaEventSourceConfig"`
	// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#source_access_configurations LambdaEventSourceMapping#source_access_configurations}
	SourceAccessConfigurations interface{} `field:"optional" json:"sourceAccessConfigurations" yaml:"sourceAccessConfigurations"`
	// The position in a stream from which to start reading.
	//
	// Required for Amazon Kinesis and Amazon DynamoDB.
	//   +  *LATEST* - Read only new records.
	//   +  *TRIM_HORIZON* - Process all available records.
	//   +  *AT_TIMESTAMP* - Specify a time from which to start reading records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#starting_position LambdaEventSourceMapping#starting_position}
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// With ``StartingPosition`` set to ``AT_TIMESTAMP``, the time from which to start reading, in Unix time seconds.
	//
	// ``StartingPositionTimestamp`` cannot be in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#starting_position_timestamp LambdaEventSourceMapping#starting_position_timestamp}
	StartingPositionTimestamp *float64 `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// A list of tags to add to the event source mapping.
	//
	// You must have the ``lambda:TagResource``, ``lambda:UntagResource``, and ``lambda:ListTags`` permissions for your [principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the CFN stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#tags LambdaEventSourceMapping#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the Kafka topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#topics LambdaEventSourceMapping#topics}
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources.
	//
	// A value of 0 seconds indicates no tumbling window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#tumbling_window_in_seconds LambdaEventSourceMapping#tumbling_window_in_seconds}
	TumblingWindowInSeconds *float64 `field:"optional" json:"tumblingWindowInSeconds" yaml:"tumblingWindowInSeconds"`
}


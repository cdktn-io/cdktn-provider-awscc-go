// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisChannelConfig struct {
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
	// The name of the channel. The name's uniqueness is scoped per AWS account and region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#channel_name KinesisChannel#channel_name}
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The ARN of the IAM role that the channel assumes to read from the source stream, deliver records to the destination, and (when enabled) write CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#service_execution_role_arn KinesisChannel#service_execution_role_arn}
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// List of stream configurations associated with the channel.
	//
	// v1 supports a single element; the list shape allows future extensibility to fan in from multiple streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#stream_configuration_list KinesisChannel#stream_configuration_list}
	StreamConfigurationList interface{} `field:"required" json:"streamConfigurationList" yaml:"streamConfigurationList"`
	// Server-side encryption configuration for data at rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#encryption_configuration KinesisChannel#encryption_configuration}
	EncryptionConfiguration *KinesisChannelEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Configuration for delivering channel operational logs. Defaults to CloudWatch Logs disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#logging_configuration KinesisChannel#logging_configuration}
	LoggingConfiguration *KinesisChannelLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Configuration for delivery to a vanilla S3 bucket destination. Exactly one of S3DestinationConfiguration and S3TablesDestinationConfiguration must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#s3_destination_configuration KinesisChannel#s3_destination_configuration}
	S3DestinationConfiguration *KinesisChannelS3DestinationConfiguration `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// Configuration for delivery to S3 Tables destinations. Exactly one of S3DestinationConfiguration and S3TablesDestinationConfiguration must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#s3_tables_destination_configuration KinesisChannel#s3_tables_destination_configuration}
	S3TablesDestinationConfiguration *KinesisChannelS3TablesDestinationConfiguration `field:"optional" json:"s3TablesDestinationConfiguration" yaml:"s3TablesDestinationConfiguration"`
	// An arbitrary set of tags (key-value pairs) to associate with the Kinesis channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#tags KinesisChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


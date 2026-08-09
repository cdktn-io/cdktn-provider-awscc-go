// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MskChannelConfig struct {
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
	// Name of the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#channel_name MskChannel#channel_name}
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Topic configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#topic_configuration_list MskChannel#topic_configuration_list}
	TopicConfigurationList interface{} `field:"required" json:"topicConfigurationList" yaml:"topicConfigurationList"`
	// The Amazon Resource Name (ARN) of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#cluster_arn MskChannel#cluster_arn}
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#encryption_configuration MskChannel#encryption_configuration}
	EncryptionConfiguration *MskChannelEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Iceberg destination configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#iceberg_destination_configuration MskChannel#iceberg_destination_configuration}
	IcebergDestinationConfiguration *MskChannelIcebergDestinationConfiguration `field:"optional" json:"icebergDestinationConfiguration" yaml:"icebergDestinationConfiguration"`
	// Log configuration details for Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#logging_info MskChannel#logging_info}
	LoggingInfo *MskChannelLoggingInfo `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// S3 destination configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#s3_destination_configuration MskChannel#s3_destination_configuration}
	S3DestinationConfiguration *MskChannelS3DestinationConfiguration `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// Tags attached to the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#tags MskChannel#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


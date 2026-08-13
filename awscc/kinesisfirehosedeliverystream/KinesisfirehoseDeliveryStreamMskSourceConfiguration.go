// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamMskSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#authentication_configuration KinesisfirehoseDeliveryStream#authentication_configuration}.
	AuthenticationConfiguration *KinesisfirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#msk_cluster_arn KinesisfirehoseDeliveryStream#msk_cluster_arn}.
	MskClusterArn *string `field:"optional" json:"mskClusterArn" yaml:"mskClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#read_from_timestamp KinesisfirehoseDeliveryStream#read_from_timestamp}.
	ReadFromTimestamp *string `field:"optional" json:"readFromTimestamp" yaml:"readFromTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#topic_name KinesisfirehoseDeliveryStream#topic_name}.
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorLogDeliveryWorkerLogDelivery struct {
	// Details about delivering logs to Amazon CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kafkaconnect_connector#cloudwatch_logs KafkaconnectConnector#cloudwatch_logs}
	CloudwatchLogs *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Details about delivering logs to Amazon Kinesis Data Firehose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kafkaconnect_connector#firehose KafkaconnectConnector#firehose}
	Firehose *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// Details about delivering logs to Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kafkaconnect_connector#s3 KafkaconnectConnector#s3}
	S3 *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryS3 `field:"optional" json:"s3" yaml:"s3"`
}


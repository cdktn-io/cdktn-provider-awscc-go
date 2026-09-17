// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorLogDeliveryWorkerLogDeliveryFirehose struct {
	// The Kinesis Data Firehose delivery stream that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kafkaconnect_connector#delivery_stream KafkaconnectConnector#delivery_stream}
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// Specifies whether the logs get sent to the specified Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kafkaconnect_connector#enabled KafkaconnectConnector#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}


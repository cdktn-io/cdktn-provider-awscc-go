// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs struct {
	// Specifies whether the logs get sent to the specified CloudWatch Logs destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kafkaconnect_connector#enabled KafkaconnectConnector#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The CloudWatch log group that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kafkaconnect_connector#log_group KafkaconnectConnector#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}


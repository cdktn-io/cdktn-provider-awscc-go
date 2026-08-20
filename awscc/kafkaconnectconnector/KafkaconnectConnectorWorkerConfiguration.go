// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorWorkerConfiguration struct {
	// The revision of the worker configuration to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kafkaconnect_connector#revision KafkaconnectConnector#revision}
	Revision *float64 `field:"optional" json:"revision" yaml:"revision"`
	// The Amazon Resource Name (ARN) of the worker configuration to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kafkaconnect_connector#worker_configuration_arn KafkaconnectConnector#worker_configuration_arn}
	WorkerConfigurationArn *string `field:"optional" json:"workerConfigurationArn" yaml:"workerConfigurationArn"`
}


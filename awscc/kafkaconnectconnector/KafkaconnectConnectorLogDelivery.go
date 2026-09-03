// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorLogDelivery struct {
	// Specifies where worker logs are delivered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kafkaconnect_connector#worker_log_delivery KafkaconnectConnector#worker_log_delivery}
	WorkerLogDelivery *KafkaconnectConnectorLogDeliveryWorkerLogDelivery `field:"optional" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}


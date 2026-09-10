// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorCapacity struct {
	// Details about auto scaling of a connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kafkaconnect_connector#auto_scaling KafkaconnectConnector#auto_scaling}
	AutoScaling *KafkaconnectConnectorCapacityAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// Details about a fixed capacity allocated to a connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kafkaconnect_connector#provisioned_capacity KafkaconnectConnector#provisioned_capacity}
	ProvisionedCapacity *KafkaconnectConnectorCapacityProvisionedCapacity `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}


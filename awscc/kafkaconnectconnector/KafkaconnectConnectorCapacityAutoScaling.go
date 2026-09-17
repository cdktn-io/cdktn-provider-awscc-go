// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorCapacityAutoScaling struct {
	// The maximum number of tasks allocated to the connector during autoscaling operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kafkaconnect_connector#max_autoscaling_task_count KafkaconnectConnector#max_autoscaling_task_count}
	MaxAutoscalingTaskCount *float64 `field:"optional" json:"maxAutoscalingTaskCount" yaml:"maxAutoscalingTaskCount"`
	// The maximum number of workers for a connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kafkaconnect_connector#max_worker_count KafkaconnectConnector#max_worker_count}
	MaxWorkerCount *float64 `field:"optional" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// Specifies how many MSK Connect Units (MCU) as the minimum scaling unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kafkaconnect_connector#mcu_count KafkaconnectConnector#mcu_count}
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
	// The minimum number of workers for a connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kafkaconnect_connector#min_worker_count KafkaconnectConnector#min_worker_count}
	MinWorkerCount *float64 `field:"optional" json:"minWorkerCount" yaml:"minWorkerCount"`
	// Information about the scale in policy of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kafkaconnect_connector#scale_in_policy KafkaconnectConnector#scale_in_policy}
	ScaleInPolicy *KafkaconnectConnectorCapacityAutoScalingScaleInPolicy `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// Information about the scale out policy of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kafkaconnect_connector#scale_out_policy KafkaconnectConnector#scale_out_policy}
	ScaleOutPolicy *KafkaconnectConnectorCapacityAutoScalingScaleOutPolicy `field:"optional" json:"scaleOutPolicy" yaml:"scaleOutPolicy"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign


type IotfleetwiseCampaignDataDestinationConfigsMqttTopicConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_campaign#execution_role_arn IotfleetwiseCampaign#execution_role_arn}.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_campaign#mqtt_topic_arn IotfleetwiseCampaign#mqtt_topic_arn}.
	MqttTopicArn *string `field:"optional" json:"mqttTopicArn" yaml:"mqttTopicArn"`
}


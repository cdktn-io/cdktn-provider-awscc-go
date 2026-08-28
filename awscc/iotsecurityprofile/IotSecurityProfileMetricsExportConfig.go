// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile


type IotSecurityProfileMetricsExportConfig struct {
	// The topic for metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_security_profile#mqtt_topic IotSecurityProfile#mqtt_topic}
	MqttTopic *string `field:"optional" json:"mqttTopic" yaml:"mqttTopic"`
	// The ARN of the role that grants permission to publish to mqtt topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_security_profile#role_arn IotSecurityProfile#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}


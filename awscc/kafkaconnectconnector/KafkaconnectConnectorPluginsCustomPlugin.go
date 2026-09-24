// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorPluginsCustomPlugin struct {
	// The Amazon Resource Name (ARN) of the custom plugin to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#custom_plugin_arn KafkaconnectConnector#custom_plugin_arn}
	CustomPluginArn *string `field:"required" json:"customPluginArn" yaml:"customPluginArn"`
	// The revision of the custom plugin to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#revision KafkaconnectConnector#revision}
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}


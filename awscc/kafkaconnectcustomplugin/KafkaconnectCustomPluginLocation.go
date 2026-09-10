// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectcustomplugin


type KafkaconnectCustomPluginLocation struct {
	// The S3 bucket Amazon Resource Name (ARN), file key, and object version of the plugin file stored in Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kafkaconnect_custom_plugin#s3_location KafkaconnectCustomPlugin#s3_location}
	S3Location *KafkaconnectCustomPluginLocationS3Location `field:"required" json:"s3Location" yaml:"s3Location"`
}


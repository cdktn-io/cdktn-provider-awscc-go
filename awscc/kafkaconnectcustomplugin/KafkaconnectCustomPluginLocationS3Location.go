// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectcustomplugin


type KafkaconnectCustomPluginLocationS3Location struct {
	// The Amazon Resource Name (ARN) of an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kafkaconnect_custom_plugin#bucket_arn KafkaconnectCustomPlugin#bucket_arn}
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The file key for an object in an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kafkaconnect_custom_plugin#file_key KafkaconnectCustomPlugin#file_key}
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// The version of an object in an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kafkaconnect_custom_plugin#object_version KafkaconnectCustomPlugin#object_version}
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}


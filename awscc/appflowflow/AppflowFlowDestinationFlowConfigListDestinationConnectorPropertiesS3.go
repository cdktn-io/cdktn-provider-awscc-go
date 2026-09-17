// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#bucket_name AppflowFlow#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#bucket_prefix AppflowFlow#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#s3_output_format_config AppflowFlow#s3_output_format_config}.
	S3OutputFormatConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfig `field:"optional" json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
}


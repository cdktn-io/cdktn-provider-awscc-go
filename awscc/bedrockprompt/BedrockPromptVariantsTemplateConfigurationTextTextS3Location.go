// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationTextTextS3Location struct {
	// A bucket in S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#bucket BedrockPrompt#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// A object key in S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#key BedrockPrompt#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The version of the the S3 object to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#version BedrockPrompt#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}


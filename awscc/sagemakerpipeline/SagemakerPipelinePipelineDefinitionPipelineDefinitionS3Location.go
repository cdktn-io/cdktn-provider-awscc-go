// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerpipeline


type SagemakerPipelinePipelineDefinitionPipelineDefinitionS3Location struct {
	// The name of the S3 bucket where the PipelineDefinition file is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_pipeline#bucket SagemakerPipeline#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The Amazon S3 ETag (a file checksum) of the PipelineDefinition file.
	//
	// If you don't specify a value, SageMaker skips ETag validation of your PipelineDefinition file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_pipeline#e_tag SagemakerPipeline#e_tag}
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// The file name of the PipelineDefinition file (Amazon S3 object name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_pipeline#key SagemakerPipeline#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// For versioning-enabled buckets, a specific version of the PipelineDefinition file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_pipeline#version SagemakerPipeline#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}


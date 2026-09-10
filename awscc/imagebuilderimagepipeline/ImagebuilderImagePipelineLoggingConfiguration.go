// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineLoggingConfiguration struct {
	// The name of the log group for image build logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_image_pipeline#image_log_group_name ImagebuilderImagePipeline#image_log_group_name}
	ImageLogGroupName *string `field:"optional" json:"imageLogGroupName" yaml:"imageLogGroupName"`
	// The name of the log group for pipeline execution logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_image_pipeline#pipeline_log_group_name ImagebuilderImagePipeline#pipeline_log_group_name}
	PipelineLogGroupName *string `field:"optional" json:"pipelineLogGroupName" yaml:"pipelineLogGroupName"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerpipeline

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerPipelineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_pipeline#pipeline_definition SagemakerPipeline#pipeline_definition}.
	PipelineDefinition *SagemakerPipelinePipelineDefinition `field:"required" json:"pipelineDefinition" yaml:"pipelineDefinition"`
	// The name of the Pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_pipeline#pipeline_name SagemakerPipeline#pipeline_name}
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// Role Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_pipeline#role_arn SagemakerPipeline#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_pipeline#parallelism_configuration SagemakerPipeline#parallelism_configuration}.
	ParallelismConfiguration *SagemakerPipelineParallelismConfiguration `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
	// The description of the Pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_pipeline#pipeline_description SagemakerPipeline#pipeline_description}
	PipelineDescription *string `field:"optional" json:"pipelineDescription" yaml:"pipelineDescription"`
	// The display name of the Pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_pipeline#pipeline_display_name SagemakerPipeline#pipeline_display_name}
	PipelineDisplayName *string `field:"optional" json:"pipelineDisplayName" yaml:"pipelineDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_pipeline#tags SagemakerPipeline#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


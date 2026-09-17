// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisepipeline

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewisePipelineConfig struct {
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
	// The list of compute nodes that form the pipeline DAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#computations IotsitewisePipeline#computations}
	Computations interface{} `field:"required" json:"computations" yaml:"computations"`
	// The name of the pipeline. Must be unique within the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#pipeline_name IotsitewisePipeline#pipeline_name}
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// The name of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#workspace_name IotsitewisePipeline#workspace_name}
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// A description of the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#description IotsitewisePipeline#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Environment variables shared across all compute nodes in the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#environment_variables IotsitewisePipeline#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#tags IotsitewisePipeline#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


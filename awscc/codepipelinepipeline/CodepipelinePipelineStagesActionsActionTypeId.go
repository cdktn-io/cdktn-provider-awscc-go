// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesActionsActionTypeId struct {
	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action.
	//
	// Valid categories are limited to one of the values below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#category CodepipelinePipeline#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// The creator of the action being called.
	//
	// There are three valid values for the Owner field in the action category section within your pipeline structure: AWS, ThirdParty, and Custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#owner CodepipelinePipeline#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The provider of the service being called by the action.
	//
	// Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of CodeDeploy, which would be specified as CodeDeploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#provider CodepipelinePipeline#provider}
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// A string that describes the action version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#version CodepipelinePipeline#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesOnSuccessConditionsRulesRuleTypeId struct {
	// A category for the provider type for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_pipeline#category CodepipelinePipeline#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The creator of the rule being called. Only AWS is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_pipeline#owner CodepipelinePipeline#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The provider of the service being called by the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_pipeline#provider CodepipelinePipeline#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// A string that describes the rule version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_pipeline#version CodepipelinePipeline#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}


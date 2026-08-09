// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesActionsEnvironmentVariables struct {
	// The name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#name CodepipelinePipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#type CodepipelinePipeline#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#value CodepipelinePipeline#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


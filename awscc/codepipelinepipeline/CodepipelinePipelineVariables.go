// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineVariables struct {
	// The value of a pipeline-level variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#default_value CodepipelinePipeline#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The description of a pipeline-level variable.
	//
	// It's used to add additional context about the variable, and not being used at time when pipeline executes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#description CodepipelinePipeline#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of a pipeline-level variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#name CodepipelinePipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStages struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#actions CodepipelinePipeline#actions}.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The name of the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#name CodepipelinePipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The method to use before stage runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#before_entry CodepipelinePipeline#before_entry}
	BeforeEntry *CodepipelinePipelineStagesBeforeEntry `field:"optional" json:"beforeEntry" yaml:"beforeEntry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#blockers CodepipelinePipeline#blockers}.
	Blockers interface{} `field:"optional" json:"blockers" yaml:"blockers"`
	// The method to use when a stage has not completed successfully.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#on_failure CodepipelinePipeline#on_failure}
	OnFailure *CodepipelinePipelineStagesOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The method to use when a stage has completed successfully.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#on_success CodepipelinePipeline#on_success}
	OnSuccess *CodepipelinePipelineStagesOnSuccess `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}


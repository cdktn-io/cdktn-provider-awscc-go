// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesActionsOutputArtifacts struct {
	// The files that you want to associate with the output artifact that will be exported from the compute action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_pipeline#files CodepipelinePipeline#files}
	Files *[]*string `field:"optional" json:"files" yaml:"files"`
	// The name of the output of an artifact, such as "My App".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_pipeline#name CodepipelinePipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


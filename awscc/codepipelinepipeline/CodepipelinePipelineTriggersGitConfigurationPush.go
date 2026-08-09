// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineTriggersGitConfigurationPush struct {
	// The Git repository branches specified as filter criteria to start the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#branches CodepipelinePipeline#branches}
	Branches *CodepipelinePipelineTriggersGitConfigurationPushBranches `field:"optional" json:"branches" yaml:"branches"`
	// The Git repository file paths specified as filter criteria to start the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#file_paths CodepipelinePipeline#file_paths}
	FilePaths *CodepipelinePipelineTriggersGitConfigurationPushFilePaths `field:"optional" json:"filePaths" yaml:"filePaths"`
	// The Git tags specified as filter criteria for whether a Git tag repository event will start the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#tags CodepipelinePipeline#tags}
	Tags *CodepipelinePipelineTriggersGitConfigurationPushTags `field:"optional" json:"tags" yaml:"tags"`
}


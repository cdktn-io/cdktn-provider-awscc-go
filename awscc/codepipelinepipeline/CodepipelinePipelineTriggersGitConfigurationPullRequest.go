// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineTriggersGitConfigurationPullRequest struct {
	// The Git repository branches specified as filter criteria to start the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_pipeline#branches CodepipelinePipeline#branches}
	Branches *CodepipelinePipelineTriggersGitConfigurationPullRequestBranches `field:"optional" json:"branches" yaml:"branches"`
	// The field that specifies which pull request events to filter on (opened, updated, closed) for the trigger configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_pipeline#events CodepipelinePipeline#events}
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// The Git repository file paths specified as filter criteria to start the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_pipeline#file_paths CodepipelinePipeline#file_paths}
	FilePaths *CodepipelinePipelineTriggersGitConfigurationPullRequestFilePaths `field:"optional" json:"filePaths" yaml:"filePaths"`
}


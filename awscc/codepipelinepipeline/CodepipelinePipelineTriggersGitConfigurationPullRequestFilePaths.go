// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineTriggersGitConfigurationPullRequestFilePaths struct {
	// The list of patterns of Git repository file paths that, when a commit is pushed, are to be excluded from starting the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#excludes CodepipelinePipeline#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// The list of patterns of Git repository file paths that, when a commit is pushed, are to be included as criteria that starts the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#includes CodepipelinePipeline#includes}
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}


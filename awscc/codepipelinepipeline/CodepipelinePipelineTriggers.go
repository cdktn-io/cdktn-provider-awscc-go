// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineTriggers struct {
	// A type of trigger configuration for Git-based source actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#git_configuration CodepipelinePipeline#git_configuration}
	GitConfiguration *CodepipelinePipelineTriggersGitConfiguration `field:"optional" json:"gitConfiguration" yaml:"gitConfiguration"`
	// The source provider for the event, such as connections configured for a repository with Git tags, for the specified trigger configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#provider_type CodepipelinePipeline#provider_type}
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
}


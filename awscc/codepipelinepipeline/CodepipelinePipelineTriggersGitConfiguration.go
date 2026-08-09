// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineTriggersGitConfiguration struct {
	// The field where the repository event that will start the pipeline is specified as pull requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#pull_request CodepipelinePipeline#pull_request}
	PullRequest interface{} `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// The field where the repository event that will start the pipeline, such as pushing Git tags, is specified with details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#push CodepipelinePipeline#push}
	Push interface{} `field:"optional" json:"push" yaml:"push"`
	// The name of the pipeline source action where the trigger configuration, such as Git tags, is specified.
	//
	// The trigger configuration will start the pipeline upon the specified change only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#source_action_name CodepipelinePipeline#source_action_name}
	SourceActionName *string `field:"optional" json:"sourceActionName" yaml:"sourceActionName"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesOnFailure struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_pipeline#conditions CodepipelinePipeline#conditions}.
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The specified result for when the failure conditions are met, such as rolling back the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_pipeline#result CodepipelinePipeline#result}
	Result *string `field:"optional" json:"result" yaml:"result"`
	// The configuration that specifies the retry configuration for a stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_pipeline#retry_configuration CodepipelinePipeline#retry_configuration}
	RetryConfiguration *CodepipelinePipelineStagesOnFailureRetryConfiguration `field:"optional" json:"retryConfiguration" yaml:"retryConfiguration"`
}


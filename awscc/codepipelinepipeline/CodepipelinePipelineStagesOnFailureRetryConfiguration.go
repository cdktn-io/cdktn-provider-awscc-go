// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesOnFailureRetryConfiguration struct {
	// The specified retry mode type for the given stage.
	//
	// FAILED_ACTIONS will retry only the failed actions. ALL_ACTIONS will retry both failed and successful
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codepipeline_pipeline#retry_mode CodepipelinePipeline#retry_mode}
	RetryMode *string `field:"optional" json:"retryMode" yaml:"retryMode"`
}


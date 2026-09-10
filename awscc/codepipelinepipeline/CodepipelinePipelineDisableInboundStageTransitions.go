// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineDisableInboundStageTransitions struct {
	// The reason given to the user that a stage is disabled, such as waiting for manual approval or manual tests.
	//
	// This message is displayed in the pipeline console UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_pipeline#reason CodepipelinePipeline#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The name of the stage where you want to disable the inbound or outbound transition of artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_pipeline#stage_name CodepipelinePipeline#stage_name}
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}


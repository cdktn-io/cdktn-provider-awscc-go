// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesOnFailureConditions struct {
	// The specified result for when the failure conditions are met, such as rolling back the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codepipeline_pipeline#result CodepipelinePipeline#result}
	Result *string `field:"optional" json:"result" yaml:"result"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codepipeline_pipeline#rules CodepipelinePipeline#rules}.
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}


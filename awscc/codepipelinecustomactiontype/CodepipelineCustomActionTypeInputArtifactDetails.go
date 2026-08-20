// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinecustomactiontype


type CodepipelineCustomActionTypeInputArtifactDetails struct {
	// The maximum number of artifacts allowed for the action type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#maximum_count CodepipelineCustomActionType#maximum_count}
	MaximumCount *float64 `field:"required" json:"maximumCount" yaml:"maximumCount"`
	// The minimum number of artifacts allowed for the action type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#minimum_count CodepipelineCustomActionType#minimum_count}
	MinimumCount *float64 `field:"required" json:"minimumCount" yaml:"minimumCount"`
}


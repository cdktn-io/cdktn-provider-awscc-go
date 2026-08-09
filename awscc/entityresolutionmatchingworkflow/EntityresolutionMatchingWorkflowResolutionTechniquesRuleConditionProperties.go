// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionProperties struct {
	// Configuration for matching behavior within rule condition properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#matching_config EntityresolutionMatchingWorkflow#matching_config}
	MatchingConfig *EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesMatchingConfig `field:"optional" json:"matchingConfig" yaml:"matchingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#rules EntityresolutionMatchingWorkflow#rules}.
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}


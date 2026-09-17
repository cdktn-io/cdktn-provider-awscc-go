// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowResolutionTechniques struct {
	// Enables the workflow to use real-time matching. Can only be set on creation for RULE_MATCHING workflows that define RuleConditionProperties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#enable_real_time_matching EntityresolutionMatchingWorkflow#enable_real_time_matching}
	EnableRealTimeMatching interface{} `field:"optional" json:"enableRealTimeMatching" yaml:"enableRealTimeMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#provider_properties EntityresolutionMatchingWorkflow#provider_properties}.
	ProviderProperties *EntityresolutionMatchingWorkflowResolutionTechniquesProviderProperties `field:"optional" json:"providerProperties" yaml:"providerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#resolution_type EntityresolutionMatchingWorkflow#resolution_type}.
	ResolutionType *string `field:"optional" json:"resolutionType" yaml:"resolutionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#rule_based_properties EntityresolutionMatchingWorkflow#rule_based_properties}.
	RuleBasedProperties *EntityresolutionMatchingWorkflowResolutionTechniquesRuleBasedProperties `field:"optional" json:"ruleBasedProperties" yaml:"ruleBasedProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#rule_condition_properties EntityresolutionMatchingWorkflow#rule_condition_properties}.
	RuleConditionProperties *EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionProperties `field:"optional" json:"ruleConditionProperties" yaml:"ruleConditionProperties"`
}


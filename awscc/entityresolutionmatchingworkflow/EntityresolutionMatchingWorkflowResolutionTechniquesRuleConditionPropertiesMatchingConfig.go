// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesMatchingConfig struct {
	// Enables transitive matching to process records across all rule levels and connect unmatched records to existing match groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/entityresolution_matching_workflow#enable_transitive_matching EntityresolutionMatchingWorkflow#enable_transitive_matching}
	EnableTransitiveMatching interface{} `field:"optional" json:"enableTransitiveMatching" yaml:"enableTransitiveMatching"`
}


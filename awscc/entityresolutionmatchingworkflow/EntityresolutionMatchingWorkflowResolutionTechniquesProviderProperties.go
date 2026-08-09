// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowResolutionTechniquesProviderProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#intermediate_source_configuration EntityresolutionMatchingWorkflow#intermediate_source_configuration}.
	IntermediateSourceConfiguration *EntityresolutionMatchingWorkflowResolutionTechniquesProviderPropertiesIntermediateSourceConfiguration `field:"optional" json:"intermediateSourceConfiguration" yaml:"intermediateSourceConfiguration"`
	// Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#provider_configuration EntityresolutionMatchingWorkflow#provider_configuration}
	ProviderConfiguration *map[string]*string `field:"optional" json:"providerConfiguration" yaml:"providerConfiguration"`
	// Arn of the Provider service being used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#provider_service_arn EntityresolutionMatchingWorkflow#provider_service_arn}
	ProviderServiceArn *string `field:"optional" json:"providerServiceArn" yaml:"providerServiceArn"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionidmappingworkflow


type EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_id_mapping_workflow#intermediate_source_configuration EntityresolutionIdMappingWorkflow#intermediate_source_configuration}.
	IntermediateSourceConfiguration *EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfiguration `field:"optional" json:"intermediateSourceConfiguration" yaml:"intermediateSourceConfiguration"`
	// Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_id_mapping_workflow#provider_configuration EntityresolutionIdMappingWorkflow#provider_configuration}
	ProviderConfiguration *map[string]*string `field:"optional" json:"providerConfiguration" yaml:"providerConfiguration"`
	// Arn of the Provider Service being used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_id_mapping_workflow#provider_service_arn EntityresolutionIdMappingWorkflow#provider_service_arn}
	ProviderServiceArn *string `field:"optional" json:"providerServiceArn" yaml:"providerServiceArn"`
}


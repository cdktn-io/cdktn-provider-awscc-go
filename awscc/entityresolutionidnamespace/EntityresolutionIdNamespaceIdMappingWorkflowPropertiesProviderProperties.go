// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionidnamespace


type EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderProperties struct {
	// Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#provider_configuration EntityresolutionIdNamespace#provider_configuration}
	ProviderConfiguration *map[string]*string `field:"optional" json:"providerConfiguration" yaml:"providerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#provider_service_arn EntityresolutionIdNamespace#provider_service_arn}.
	ProviderServiceArn *string `field:"optional" json:"providerServiceArn" yaml:"providerServiceArn"`
}


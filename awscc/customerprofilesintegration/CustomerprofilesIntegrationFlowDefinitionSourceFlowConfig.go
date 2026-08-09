// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_integration#connector_profile_name CustomerprofilesIntegration#connector_profile_name}.
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_integration#connector_type CustomerprofilesIntegration#connector_type}.
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_integration#incremental_pull_config CustomerprofilesIntegration#incremental_pull_config}.
	IncrementalPullConfig *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigIncrementalPullConfig `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_integration#source_connector_properties CustomerprofilesIntegration#source_connector_properties}.
	SourceConnectorProperties *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorProperties `field:"optional" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
}


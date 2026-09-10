// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector


type Inspectorv2ConnectorProviderConfigurationAzureScopeConfiguration struct {
	// Defines the scope of Azure resources to monitor for a specific resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/inspectorv2_connector#container_image_scanning Inspectorv2Connector#container_image_scanning}
	ContainerImageScanning *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationContainerImageScanning `field:"optional" json:"containerImageScanning" yaml:"containerImageScanning"`
	// Defines the scope of Azure resources to monitor for a specific resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/inspectorv2_connector#serverless_scanning Inspectorv2Connector#serverless_scanning}
	ServerlessScanning *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanning `field:"optional" json:"serverlessScanning" yaml:"serverlessScanning"`
	// Defines the scope of Azure resources to monitor for a specific resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/inspectorv2_connector#vm_scanning Inspectorv2Connector#vm_scanning}
	VmScanning *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanning `field:"optional" json:"vmScanning" yaml:"vmScanning"`
}


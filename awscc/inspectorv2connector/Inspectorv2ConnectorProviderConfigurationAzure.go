// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector


type Inspectorv2ConnectorProviderConfigurationAzure struct {
	// The ARN of the AWS Config connector used for Azure resource discovery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_connector#aws_config_connector_arn Inspectorv2Connector#aws_config_connector_arn}
	AwsConfigConnectorArn *string `field:"required" json:"awsConfigConnectorArn" yaml:"awsConfigConnectorArn"`
	// List of Azure regions to scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_connector#azure_regions Inspectorv2Connector#azure_regions}
	AzureRegions *[]*string `field:"required" json:"azureRegions" yaml:"azureRegions"`
	// Defines which resource types to scan and at what scope level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_connector#scope_configuration Inspectorv2Connector#scope_configuration}
	ScopeConfiguration *Inspectorv2ConnectorProviderConfigurationAzureScopeConfiguration `field:"required" json:"scopeConfiguration" yaml:"scopeConfiguration"`
	// Whether to automatically install the VM scanner. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_connector#auto_install_vm_scanner Inspectorv2Connector#auto_install_vm_scanner}
	AutoInstallVmScanner interface{} `field:"optional" json:"autoInstallVmScanner" yaml:"autoInstallVmScanner"`
}


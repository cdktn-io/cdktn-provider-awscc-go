// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnector


type SecurityhubConnectorProviderNameAzure struct {
	// The ARN of the multi-cloud configuration connector used to establish the connection to Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_connector#aws_config_connector_arn SecurityhubConnector#aws_config_connector_arn}
	AwsConfigConnectorArn *string `field:"required" json:"awsConfigConnectorArn" yaml:"awsConfigConnectorArn"`
	// The list of Azure regions to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_connector#azure_regions SecurityhubConnector#azure_regions}
	AzureRegions *[]*string `field:"required" json:"azureRegions" yaml:"azureRegions"`
	// The scope configuration that defines which Azure resources are monitored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_connector#scope_configuration SecurityhubConnector#scope_configuration}
	ScopeConfiguration *SecurityhubConnectorProviderNameAzureScopeConfiguration `field:"required" json:"scopeConfiguration" yaml:"scopeConfiguration"`
}


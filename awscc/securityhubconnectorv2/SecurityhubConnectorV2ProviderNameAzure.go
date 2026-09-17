// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnectorv2


type SecurityhubConnectorV2ProviderNameAzure struct {
	// The ARN of the AWS Config connector used for the Azure integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector_v2#aws_config_connector_arn SecurityhubConnectorV2#aws_config_connector_arn}
	AwsConfigConnectorArn *string `field:"optional" json:"awsConfigConnectorArn" yaml:"awsConfigConnectorArn"`
	// The list of Azure regions to include in the connector scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector_v2#azure_regions SecurityhubConnectorV2#azure_regions}
	AzureRegions *[]*string `field:"optional" json:"azureRegions" yaml:"azureRegions"`
	// The scope configuration for an Azure connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector_v2#scope_configuration SecurityhubConnectorV2#scope_configuration}
	ScopeConfiguration *SecurityhubConnectorV2ProviderNameAzureScopeConfiguration `field:"optional" json:"scopeConfiguration" yaml:"scopeConfiguration"`
}


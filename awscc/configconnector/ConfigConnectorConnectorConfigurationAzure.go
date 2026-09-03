// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconnector


type ConfigConnectorConnectorConfigurationAzure struct {
	// The Azure client (application) identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/config_connector#client_identifier ConfigConnector#client_identifier}
	ClientIdentifier *string `field:"optional" json:"clientIdentifier" yaml:"clientIdentifier"`
	// The Azure tenant identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/config_connector#tenant_identifier ConfigConnector#tenant_identifier}
	TenantIdentifier *string `field:"optional" json:"tenantIdentifier" yaml:"tenantIdentifier"`
}


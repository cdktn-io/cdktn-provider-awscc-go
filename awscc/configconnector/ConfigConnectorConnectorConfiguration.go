// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconnector


type ConfigConnectorConnectorConfiguration struct {
	// The configuration for connecting to Microsoft Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/config_connector#azure ConfigConnector#azure}
	Azure *ConfigConnectorConnectorConfigurationAzure `field:"optional" json:"azure" yaml:"azure"`
}


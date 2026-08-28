// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnector


type AppflowConnectorConnectorProvisioningConfig struct {
	// Contains information about the configuration of the lambda which is being registered as the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_connector#lambda AppflowConnector#lambda}
	Lambda *AppflowConnectorConnectorProvisioningConfigLambda `field:"optional" json:"lambda" yaml:"lambda"`
}


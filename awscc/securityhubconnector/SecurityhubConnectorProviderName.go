// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnector


type SecurityhubConnectorProviderName struct {
	// The configuration for connecting to an Azure environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector#azure SecurityhubConnector#azure}
	Azure *SecurityhubConnectorProviderNameAzure `field:"required" json:"azure" yaml:"azure"`
}


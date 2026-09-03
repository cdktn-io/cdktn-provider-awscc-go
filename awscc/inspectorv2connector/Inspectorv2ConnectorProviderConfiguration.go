// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector


type Inspectorv2ConnectorProviderConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_connector#azure Inspectorv2Connector#azure}.
	Azure *Inspectorv2ConnectorProviderConfigurationAzure `field:"required" json:"azure" yaml:"azure"`
}


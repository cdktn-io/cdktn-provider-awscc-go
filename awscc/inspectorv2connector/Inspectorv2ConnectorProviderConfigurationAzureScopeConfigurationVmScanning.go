// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector


type Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/inspectorv2_connector#scope_type Inspectorv2Connector#scope_type}.
	ScopeType *string `field:"optional" json:"scopeType" yaml:"scopeType"`
	// List of subscription IDs. Empty for TENANT scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/inspectorv2_connector#scope_values Inspectorv2Connector#scope_values}
	ScopeValues *[]*string `field:"optional" json:"scopeValues" yaml:"scopeValues"`
}


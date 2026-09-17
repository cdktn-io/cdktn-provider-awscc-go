// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnectorv2


type SecurityhubConnectorV2ProviderNameAzureScopeConfiguration struct {
	// The scope type for the Azure connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector_v2#scope_type SecurityhubConnectorV2#scope_type}
	ScopeType *string `field:"optional" json:"scopeType" yaml:"scopeType"`
	// The list of scope values for the Azure connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector_v2#scope_values SecurityhubConnectorV2#scope_values}
	ScopeValues *[]*string `field:"optional" json:"scopeValues" yaml:"scopeValues"`
}


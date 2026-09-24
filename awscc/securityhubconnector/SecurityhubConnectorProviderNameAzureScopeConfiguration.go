// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnector


type SecurityhubConnectorProviderNameAzureScopeConfiguration struct {
	// The type of scope. Valid values are ``tenant`` and ``subscription``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_connector#scope_type SecurityhubConnector#scope_type}
	ScopeType *string `field:"required" json:"scopeType" yaml:"scopeType"`
	// The list of scope values, such as subscription IDs, when the scope type is ``subscription``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_connector#scope_values SecurityhubConnector#scope_values}
	ScopeValues *[]*string `field:"optional" json:"scopeValues" yaml:"scopeValues"`
}


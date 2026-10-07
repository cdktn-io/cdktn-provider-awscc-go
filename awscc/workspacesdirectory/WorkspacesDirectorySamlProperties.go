// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectorySamlProperties struct {
	// The relay state parameter name supported by the SAML 2.0 identity provider (IdP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#relay_state_parameter_name WorkspacesDirectory#relay_state_parameter_name}
	RelayStateParameterName *string `field:"optional" json:"relayStateParameterName" yaml:"relayStateParameterName"`
	// Indicates the status of SAML 2.0 authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#status WorkspacesDirectory#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The SAML 2.0 identity provider (IdP) user access URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#user_access_url WorkspacesDirectory#user_access_url}
	UserAccessUrl *string `field:"optional" json:"userAccessUrl" yaml:"userAccessUrl"`
}


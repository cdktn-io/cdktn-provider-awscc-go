// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryWorkspaceAccessPropertiesAccessEndpointConfig struct {
	// Indicates a list of access endpoints associated with this directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#access_endpoints WorkspacesDirectory#access_endpoints}
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// Indicates a list of protocols that fallback to using the public Internet when streaming over a VPC endpoint is not available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#internet_fallback_protocols WorkspacesDirectory#internet_fallback_protocols}
	InternetFallbackProtocols *[]*string `field:"optional" json:"internetFallbackProtocols" yaml:"internetFallbackProtocols"`
}


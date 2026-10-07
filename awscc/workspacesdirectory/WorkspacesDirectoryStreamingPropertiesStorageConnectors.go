// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryStreamingPropertiesStorageConnectors struct {
	// The type of connector used to save user files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#connector_type WorkspacesDirectory#connector_type}
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// Indicates if the storage connector is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#status WorkspacesDirectory#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


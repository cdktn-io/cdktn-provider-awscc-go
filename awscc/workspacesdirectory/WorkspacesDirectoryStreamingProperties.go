// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryStreamingProperties struct {
	// Describes the Global Accelerator for directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#global_accelerator WorkspacesDirectory#global_accelerator}
	GlobalAccelerator *WorkspacesDirectoryStreamingPropertiesGlobalAccelerator `field:"optional" json:"globalAccelerator" yaml:"globalAccelerator"`
	// Indicates the storage connector used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#storage_connectors WorkspacesDirectory#storage_connectors}
	StorageConnectors interface{} `field:"optional" json:"storageConnectors" yaml:"storageConnectors"`
	// Indicates the type of preferred protocol for the streaming experience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#streaming_experience_preferred_protocol WorkspacesDirectory#streaming_experience_preferred_protocol}
	StreamingExperiencePreferredProtocol *string `field:"optional" json:"streamingExperiencePreferredProtocol" yaml:"streamingExperiencePreferredProtocol"`
	// Indicates the permission settings associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#user_settings WorkspacesDirectory#user_settings}
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}


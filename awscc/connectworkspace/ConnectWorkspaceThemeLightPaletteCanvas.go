// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspaceThemeLightPaletteCanvas struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_workspace#active_background ConnectWorkspace#active_background}.
	ActiveBackground *string `field:"optional" json:"activeBackground" yaml:"activeBackground"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_workspace#container_background ConnectWorkspace#container_background}.
	ContainerBackground *string `field:"optional" json:"containerBackground" yaml:"containerBackground"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_workspace#page_background ConnectWorkspace#page_background}.
	PageBackground *string `field:"optional" json:"pageBackground" yaml:"pageBackground"`
}


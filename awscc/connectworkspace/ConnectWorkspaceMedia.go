// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspaceMedia struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#source ConnectWorkspace#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The type of media.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#type ConnectWorkspace#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


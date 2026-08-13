// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebsessionlogger


type WorkspaceswebSessionLoggerEventFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/workspacesweb_session_logger#all WorkspaceswebSessionLogger#all}.
	All *string `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/workspacesweb_session_logger#include WorkspaceswebSessionLogger#include}.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}


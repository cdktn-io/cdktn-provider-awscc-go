// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebsessionlogger


type WorkspaceswebSessionLoggerLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesweb_session_logger#s3 WorkspaceswebSessionLogger#s3}.
	S3 *WorkspaceswebSessionLoggerLogConfigurationS3 `field:"optional" json:"s3" yaml:"s3"`
}


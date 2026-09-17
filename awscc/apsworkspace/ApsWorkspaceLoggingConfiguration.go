// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsworkspace


type ApsWorkspaceLoggingConfiguration struct {
	// CloudWatch log group ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/aps_workspace#log_group_arn ApsWorkspace#log_group_arn}
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}


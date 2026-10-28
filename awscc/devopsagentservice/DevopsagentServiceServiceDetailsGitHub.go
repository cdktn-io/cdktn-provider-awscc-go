// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsGitHub struct {
	// The token verifies an existing DevOps Agent GitHub App installation on github.com and is then discarded. Not supported for GitHub Enterprise Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_service#app_verification DevopsagentService#app_verification}
	AppVerification *DevopsagentServiceServiceDetailsGitHubAppVerification `field:"optional" json:"appVerification" yaml:"appVerification"`
	// The token is stored and used as the long-lived credential for GitHub API access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_service#direct_access DevopsagentService#direct_access}
	DirectAccess *DevopsagentServiceServiceDetailsGitHubDirectAccess `field:"optional" json:"directAccess" yaml:"directAccess"`
}


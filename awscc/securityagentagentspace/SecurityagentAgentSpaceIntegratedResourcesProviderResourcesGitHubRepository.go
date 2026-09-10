// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepository struct {
	// GitHub repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityagent_agent_space#name SecurityagentAgentSpace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// GitHub repository owner (user or organization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityagent_agent_space#owner SecurityagentAgentSpace#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
}


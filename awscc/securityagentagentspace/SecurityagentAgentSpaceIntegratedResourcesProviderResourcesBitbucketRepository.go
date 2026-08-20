// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepository struct {
	// Bitbucket repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityagent_agent_space#name SecurityagentAgentSpace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Bitbucket workspace slug owning the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityagent_agent_space#workspace SecurityagentAgentSpace#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}


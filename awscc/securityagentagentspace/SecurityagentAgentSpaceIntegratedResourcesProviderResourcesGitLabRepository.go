// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepository struct {
	// GitLab project name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityagent_agent_space#name SecurityagentAgentSpace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// GitLab project namespace (user, group, or subgroup path).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityagent_agent_space#namespace SecurityagentAgentSpace#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}


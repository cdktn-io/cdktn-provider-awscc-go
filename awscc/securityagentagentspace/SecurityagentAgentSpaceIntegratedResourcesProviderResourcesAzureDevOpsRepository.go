// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesAzureDevOpsRepository struct {
	// Azure DevOps repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#name SecurityagentAgentSpace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Read-only Azure DevOps organization that owns the repository, populated from service-side metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#organization SecurityagentAgentSpace#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Azure DevOps project that owns the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#project SecurityagentAgentSpace#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}


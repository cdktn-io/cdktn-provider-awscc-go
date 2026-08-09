// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilities struct {
	// Enables Code Review in the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityagent_agent_space#leave_comments SecurityagentAgentSpace#leave_comments}
	LeaveComments interface{} `field:"optional" json:"leaveComments" yaml:"leaveComments"`
	// Enables creation of pull requests with automated fixes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityagent_agent_space#remediate_code SecurityagentAgentSpace#remediate_code}
	RemediateCode interface{} `field:"optional" json:"remediateCode" yaml:"remediateCode"`
}


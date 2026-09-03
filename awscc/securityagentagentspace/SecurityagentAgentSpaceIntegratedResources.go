// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResources struct {
	// Unique identifier of the Provider Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityagent_agent_space#integration SecurityagentAgentSpace#integration}
	Integration *string `field:"optional" json:"integration" yaml:"integration"`
	// List of selected Resources from the Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityagent_agent_space#provider_resources SecurityagentAgentSpace#provider_resources}
	ProviderResources interface{} `field:"optional" json:"providerResources" yaml:"providerResources"`
}


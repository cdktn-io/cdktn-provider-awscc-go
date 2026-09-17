// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorepolicy


type BedrockagentcorePolicyDefinition struct {
	// A Cedar policy statement within the AgentCore Policy system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_policy#cedar BedrockagentcorePolicy#cedar}
	Cedar *BedrockagentcorePolicyDefinitionCedar `field:"optional" json:"cedar" yaml:"cedar"`
	// A policy statement within the AgentCore Policy system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_policy#policy BedrockagentcorePolicy#policy}
	Policy *BedrockagentcorePolicyDefinitionPolicy `field:"optional" json:"policy" yaml:"policy"`
}


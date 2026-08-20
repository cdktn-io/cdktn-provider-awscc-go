// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorepolicy


type BedrockagentcorePolicyDefinitionPolicy struct {
	// The policy statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_policy#statement BedrockagentcorePolicy#statement}
	Statement *string `field:"optional" json:"statement" yaml:"statement"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule


type BedrockagentcoreGatewayRuleConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_gateway_rule#match_paths BedrockagentcoreGatewayRule#match_paths}.
	MatchPaths *BedrockagentcoreGatewayRuleConditionsMatchPaths `field:"optional" json:"matchPaths" yaml:"matchPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_gateway_rule#match_principals BedrockagentcoreGatewayRule#match_principals}.
	MatchPrincipals *BedrockagentcoreGatewayRuleConditionsMatchPrincipals `field:"optional" json:"matchPrincipals" yaml:"matchPrincipals"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule


type BedrockagentcoreGatewayRuleActionsRouteToTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway_rule#static_route BedrockagentcoreGatewayRule#static_route}.
	StaticRoute *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRoute `field:"optional" json:"staticRoute" yaml:"staticRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway_rule#weighted_route BedrockagentcoreGatewayRule#weighted_route}.
	WeightedRoute *BedrockagentcoreGatewayRuleActionsRouteToTargetWeightedRoute `field:"optional" json:"weightedRoute" yaml:"weightedRoute"`
}


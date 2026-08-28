// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule


type BedrockagentcoreGatewayRuleActionsConfigurationBundle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#static_override BedrockagentcoreGatewayRule#static_override}.
	StaticOverride *BedrockagentcoreGatewayRuleActionsConfigurationBundleStaticOverride `field:"optional" json:"staticOverride" yaml:"staticOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#weighted_override BedrockagentcoreGatewayRule#weighted_override}.
	WeightedOverride *BedrockagentcoreGatewayRuleActionsConfigurationBundleWeightedOverride `field:"optional" json:"weightedOverride" yaml:"weightedOverride"`
}


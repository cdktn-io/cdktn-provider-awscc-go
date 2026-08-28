// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule


type BedrockagentcoreGatewayRuleActionsRouteToTargetWeightedRouteTrafficSplit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#description BedrockagentcoreGatewayRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#metadata BedrockagentcoreGatewayRule#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#name BedrockagentcoreGatewayRule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#target_name BedrockagentcoreGatewayRule#target_name}.
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#weight BedrockagentcoreGatewayRule#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


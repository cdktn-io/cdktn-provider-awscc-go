// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule


type BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#fallback BedrockagentcoreGatewayRule#fallback}.
	Fallback *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallback `field:"optional" json:"fallback" yaml:"fallback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#model BedrockagentcoreGatewayRule#model}.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#target_name BedrockagentcoreGatewayRule#target_name}.
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
}


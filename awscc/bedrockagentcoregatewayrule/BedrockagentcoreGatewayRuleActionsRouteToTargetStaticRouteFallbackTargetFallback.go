// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule


type BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallback struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#first_byte_timeout_milliseconds BedrockagentcoreGatewayRule#first_byte_timeout_milliseconds}.
	FirstByteTimeoutMilliseconds *float64 `field:"optional" json:"firstByteTimeoutMilliseconds" yaml:"firstByteTimeoutMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#model BedrockagentcoreGatewayRule#model}.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#target_name BedrockagentcoreGatewayRule#target_name}.
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#trigger_on BedrockagentcoreGatewayRule#trigger_on}.
	TriggerOn *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackTriggerOn `field:"optional" json:"triggerOn" yaml:"triggerOn"`
}


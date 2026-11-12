// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule


type BedrockagentcoreGatewayRuleActionsRouteToTargetWeightedRouteTrafficSplitFallback struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#circuit_breaker BedrockagentcoreGatewayRule#circuit_breaker}.
	CircuitBreaker *BedrockagentcoreGatewayRuleActionsRouteToTargetWeightedRouteTrafficSplitFallbackCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#first_byte_timeout_milliseconds BedrockagentcoreGatewayRule#first_byte_timeout_milliseconds}.
	FirstByteTimeoutMilliseconds *float64 `field:"optional" json:"firstByteTimeoutMilliseconds" yaml:"firstByteTimeoutMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#target_fallback BedrockagentcoreGatewayRule#target_fallback}.
	TargetFallback interface{} `field:"optional" json:"targetFallback" yaml:"targetFallback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/bedrockagentcore_gateway_rule#trigger_on BedrockagentcoreGatewayRule#trigger_on}.
	TriggerOn *BedrockagentcoreGatewayRuleActionsRouteToTargetWeightedRouteTrafficSplitFallbackTriggerOn `field:"optional" json:"triggerOn" yaml:"triggerOn"`
}


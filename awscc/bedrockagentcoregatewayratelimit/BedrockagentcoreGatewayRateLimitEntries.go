// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayratelimit


type BedrockagentcoreGatewayRateLimitEntries struct {
	// Map of dimension name to dimension value for a rule entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_gateway_rate_limit#dimensions BedrockagentcoreGatewayRateLimit#dimensions}
	Dimensions *map[string]*string `field:"required" json:"dimensions" yaml:"dimensions"`
	// Connection rate limits (per second only). Limited to 1 entry for now. — P2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_gateway_rate_limit#connections BedrockagentcoreGatewayRateLimit#connections}
	Connections interface{} `field:"optional" json:"connections" yaml:"connections"`
	// Request rate limits (RPS or RPM). Limited to 1 entry for now.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_gateway_rate_limit#requests BedrockagentcoreGatewayRateLimit#requests}
	Requests interface{} `field:"optional" json:"requests" yaml:"requests"`
	// Token rate limits (TPM). Limited to 1 entry for now. — P1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_gateway_rate_limit#tokens BedrockagentcoreGatewayRateLimit#tokens}
	Tokens interface{} `field:"optional" json:"tokens" yaml:"tokens"`
}


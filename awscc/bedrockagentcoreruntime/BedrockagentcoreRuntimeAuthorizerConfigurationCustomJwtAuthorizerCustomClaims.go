// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerCustomClaims struct {
	// The value or values in the custom claim to match and relationship of match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#authorizing_claim_match_value BedrockagentcoreRuntime#authorizing_claim_match_value}
	AuthorizingClaimMatchValue *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerCustomClaimsAuthorizingClaimMatchValue `field:"optional" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
	// The name of the custom claim to validate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#inbound_token_claim_name BedrockagentcoreRuntime#inbound_token_claim_name}
	InboundTokenClaimName *string `field:"optional" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// Token claim data type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#inbound_token_claim_value_type BedrockagentcoreRuntime#inbound_token_claim_value_type}
	InboundTokenClaimValueType *string `field:"optional" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
}


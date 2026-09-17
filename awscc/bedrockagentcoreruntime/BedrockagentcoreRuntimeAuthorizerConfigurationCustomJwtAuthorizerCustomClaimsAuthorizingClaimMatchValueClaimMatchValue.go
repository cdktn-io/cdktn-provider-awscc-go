// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerCustomClaimsAuthorizingClaimMatchValueClaimMatchValue struct {
	// The string value to match for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#match_value_string BedrockagentcoreRuntime#match_value_string}
	MatchValueString *string `field:"optional" json:"matchValueString" yaml:"matchValueString"`
	// The list of strings to check for a match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#match_value_string_list BedrockagentcoreRuntime#match_value_string_list}
	MatchValueStringList *[]*string `field:"optional" json:"matchValueStringList" yaml:"matchValueStringList"`
}


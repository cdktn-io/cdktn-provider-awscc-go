// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAuthorizerConfiguration struct {
	// Configuration for custom JWT authorizer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_runtime#custom_jwt_authorizer BedrockagentcoreRuntime#custom_jwt_authorizer}
	CustomJwtAuthorizer *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizer `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}


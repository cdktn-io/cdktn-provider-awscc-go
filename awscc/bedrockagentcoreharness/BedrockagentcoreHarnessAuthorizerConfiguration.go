// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessAuthorizerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#custom_jwt_authorizer BedrockagentcoreHarness#custom_jwt_authorizer}.
	CustomJwtAuthorizer *BedrockagentcoreHarnessAuthorizerConfigurationCustomJwtAuthorizer `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessAuthorizerConfigurationCustomJwtAuthorizer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#allowed_audience BedrockagentcoreHarness#allowed_audience}.
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#allowed_clients BedrockagentcoreHarness#allowed_clients}.
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#allowed_scopes BedrockagentcoreHarness#allowed_scopes}.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#custom_claims BedrockagentcoreHarness#custom_claims}.
	CustomClaims interface{} `field:"optional" json:"customClaims" yaml:"customClaims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#discovery_url BedrockagentcoreHarness#discovery_url}.
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
}


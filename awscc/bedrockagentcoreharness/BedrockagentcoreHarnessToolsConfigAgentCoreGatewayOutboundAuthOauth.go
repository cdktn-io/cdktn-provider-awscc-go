// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessToolsConfigAgentCoreGatewayOutboundAuthOauth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#custom_parameters BedrockagentcoreHarness#custom_parameters}.
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#default_return_url BedrockagentcoreHarness#default_return_url}.
	DefaultReturnUrl *string `field:"optional" json:"defaultReturnUrl" yaml:"defaultReturnUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#grant_type BedrockagentcoreHarness#grant_type}.
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#provider_arn BedrockagentcoreHarness#provider_arn}.
	ProviderArn *string `field:"optional" json:"providerArn" yaml:"providerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#scopes BedrockagentcoreHarness#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}


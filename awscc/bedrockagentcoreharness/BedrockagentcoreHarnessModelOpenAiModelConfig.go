// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessModelOpenAiModelConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#api_key_arn BedrockagentcoreHarness#api_key_arn}.
	ApiKeyArn *string `field:"optional" json:"apiKeyArn" yaml:"apiKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#max_tokens BedrockagentcoreHarness#max_tokens}.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#model_id BedrockagentcoreHarness#model_id}.
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#temperature BedrockagentcoreHarness#temperature}.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#top_p BedrockagentcoreHarness#top_p}.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


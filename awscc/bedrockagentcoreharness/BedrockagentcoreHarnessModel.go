// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness#bedrock_model_config BedrockagentcoreHarness#bedrock_model_config}.
	BedrockModelConfig *BedrockagentcoreHarnessModelBedrockModelConfig `field:"optional" json:"bedrockModelConfig" yaml:"bedrockModelConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness#gemini_model_config BedrockagentcoreHarness#gemini_model_config}.
	GeminiModelConfig *BedrockagentcoreHarnessModelGeminiModelConfig `field:"optional" json:"geminiModelConfig" yaml:"geminiModelConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness#open_ai_model_config BedrockagentcoreHarness#open_ai_model_config}.
	OpenAiModelConfig *BedrockagentcoreHarnessModelOpenAiModelConfig `field:"optional" json:"openAiModelConfig" yaml:"openAiModelConfig"`
}


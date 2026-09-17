// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockintelligentpromptrouter


type BedrockIntelligentPromptRouterModels struct {
	// Arn of underlying model which are added in the Prompt Router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_intelligent_prompt_router#model_arn BedrockIntelligentPromptRouter#model_arn}
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
}


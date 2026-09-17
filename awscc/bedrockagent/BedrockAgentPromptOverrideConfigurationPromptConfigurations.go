// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentPromptOverrideConfigurationPromptConfigurations struct {
	// Additional Model Request Fields for Prompt Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#additional_model_request_fields BedrockAgent#additional_model_request_fields}
	AdditionalModelRequestFields *string `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// Base Prompt Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#base_prompt_template BedrockAgent#base_prompt_template}
	BasePromptTemplate *string `field:"optional" json:"basePromptTemplate" yaml:"basePromptTemplate"`
	// ARN or name of a Bedrock model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#foundation_model BedrockAgent#foundation_model}
	FoundationModel *string `field:"optional" json:"foundationModel" yaml:"foundationModel"`
	// Configuration for inference in prompt configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#inference_configuration BedrockAgent#inference_configuration}
	InferenceConfiguration *BedrockAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfiguration `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Creation Mode for Prompt Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#parser_mode BedrockAgent#parser_mode}
	ParserMode *string `field:"optional" json:"parserMode" yaml:"parserMode"`
	// Creation Mode for Prompt Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#prompt_creation_mode BedrockAgent#prompt_creation_mode}
	PromptCreationMode *string `field:"optional" json:"promptCreationMode" yaml:"promptCreationMode"`
	// Prompt State.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#prompt_state BedrockAgent#prompt_state}
	PromptState *string `field:"optional" json:"promptState" yaml:"promptState"`
	// Prompt Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#prompt_type BedrockAgent#prompt_type}
	PromptType *string `field:"optional" json:"promptType" yaml:"promptType"`
}


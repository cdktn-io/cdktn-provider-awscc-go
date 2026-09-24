// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariants struct {
	// Contains model-specific configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#additional_model_request_fields BedrockPrompt#additional_model_request_fields}
	AdditionalModelRequestFields *string `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// Target resource to invoke with Prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#gen_ai_resource BedrockPrompt#gen_ai_resource}
	GenAiResource *BedrockPromptVariantsGenAiResource `field:"optional" json:"genAiResource" yaml:"genAiResource"`
	// Model inference configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#inference_configuration BedrockPrompt#inference_configuration}
	InferenceConfiguration *BedrockPromptVariantsInferenceConfiguration `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// List of metadata to associate with the prompt variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#metadata BedrockPrompt#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// ARN or Id of a Bedrock Foundational Model or Inference Profile, or the ARN of a imported model, or a provisioned throughput ARN for custom models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#model_id BedrockPrompt#model_id}
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// Name for a variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#name BedrockPrompt#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Prompt template configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#template_configuration BedrockPrompt#template_configuration}
	TemplateConfiguration *BedrockPromptVariantsTemplateConfiguration `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// Prompt template type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#template_type BedrockPrompt#template_type}
	TemplateType *string `field:"optional" json:"templateType" yaml:"templateType"`
}


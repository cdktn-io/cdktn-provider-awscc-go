// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectOverrideConfigurationAudioLanguageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#generative_output_language BedrockDataAutomationProject#generative_output_language}.
	GenerativeOutputLanguage *string `field:"optional" json:"generativeOutputLanguage" yaml:"generativeOutputLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#identify_multiple_languages BedrockDataAutomationProject#identify_multiple_languages}.
	IdentifyMultipleLanguages interface{} `field:"optional" json:"identifyMultipleLanguages" yaml:"identifyMultipleLanguages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#input_languages BedrockDataAutomationProject#input_languages}.
	InputLanguages *[]*string `field:"optional" json:"inputLanguages" yaml:"inputLanguages"`
}


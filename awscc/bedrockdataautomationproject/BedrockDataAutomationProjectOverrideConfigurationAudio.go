// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectOverrideConfigurationAudio struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#language_configuration BedrockDataAutomationProject#language_configuration}.
	LanguageConfiguration *BedrockDataAutomationProjectOverrideConfigurationAudioLanguageConfiguration `field:"optional" json:"languageConfiguration" yaml:"languageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#modality_processing BedrockDataAutomationProject#modality_processing}.
	ModalityProcessing *BedrockDataAutomationProjectOverrideConfigurationAudioModalityProcessing `field:"optional" json:"modalityProcessing" yaml:"modalityProcessing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#sensitive_data_configuration BedrockDataAutomationProject#sensitive_data_configuration}.
	SensitiveDataConfiguration *BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfiguration `field:"optional" json:"sensitiveDataConfiguration" yaml:"sensitiveDataConfiguration"`
}


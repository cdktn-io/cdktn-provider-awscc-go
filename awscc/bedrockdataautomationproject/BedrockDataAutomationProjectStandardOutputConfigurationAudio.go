// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationAudio struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_automation_project#extraction BedrockDataAutomationProject#extraction}.
	Extraction *BedrockDataAutomationProjectStandardOutputConfigurationAudioExtraction `field:"optional" json:"extraction" yaml:"extraction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_automation_project#generative_field BedrockDataAutomationProject#generative_field}.
	GenerativeField *BedrockDataAutomationProjectStandardOutputConfigurationAudioGenerativeField `field:"optional" json:"generativeField" yaml:"generativeField"`
}


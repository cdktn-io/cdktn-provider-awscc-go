// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationVideo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#extraction BedrockDataAutomationProject#extraction}.
	Extraction *BedrockDataAutomationProjectStandardOutputConfigurationVideoExtraction `field:"optional" json:"extraction" yaml:"extraction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#generative_field BedrockDataAutomationProject#generative_field}.
	GenerativeField *BedrockDataAutomationProjectStandardOutputConfigurationVideoGenerativeField `field:"optional" json:"generativeField" yaml:"generativeField"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationDocument struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#extraction BedrockDataAutomationProject#extraction}.
	Extraction *BedrockDataAutomationProjectStandardOutputConfigurationDocumentExtraction `field:"optional" json:"extraction" yaml:"extraction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#generative_field BedrockDataAutomationProject#generative_field}.
	GenerativeField *BedrockDataAutomationProjectStandardOutputConfigurationDocumentGenerativeField `field:"optional" json:"generativeField" yaml:"generativeField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#output_format BedrockDataAutomationProject#output_format}.
	OutputFormat *BedrockDataAutomationProjectStandardOutputConfigurationDocumentOutputFormat `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}


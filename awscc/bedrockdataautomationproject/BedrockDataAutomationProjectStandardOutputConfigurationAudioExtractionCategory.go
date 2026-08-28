// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#state BedrockDataAutomationProject#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#type_configuration BedrockDataAutomationProject#type_configuration}.
	TypeConfiguration *BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfiguration `field:"optional" json:"typeConfiguration" yaml:"typeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#types BedrockDataAutomationProject#types}.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}


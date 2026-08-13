// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#audio BedrockDataAutomationProject#audio}.
	Audio *BedrockDataAutomationProjectStandardOutputConfigurationAudio `field:"optional" json:"audio" yaml:"audio"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#document BedrockDataAutomationProject#document}.
	Document *BedrockDataAutomationProjectStandardOutputConfigurationDocument `field:"optional" json:"document" yaml:"document"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#image BedrockDataAutomationProject#image}.
	Image *BedrockDataAutomationProjectStandardOutputConfigurationImage `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#video BedrockDataAutomationProject#video}.
	Video *BedrockDataAutomationProjectStandardOutputConfigurationVideo `field:"optional" json:"video" yaml:"video"`
}


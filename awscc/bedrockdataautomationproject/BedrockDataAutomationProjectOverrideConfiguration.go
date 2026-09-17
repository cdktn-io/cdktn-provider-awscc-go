// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectOverrideConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_automation_project#audio BedrockDataAutomationProject#audio}.
	Audio *BedrockDataAutomationProjectOverrideConfigurationAudio `field:"optional" json:"audio" yaml:"audio"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_automation_project#document BedrockDataAutomationProject#document}.
	Document *BedrockDataAutomationProjectOverrideConfigurationDocument `field:"optional" json:"document" yaml:"document"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_automation_project#image BedrockDataAutomationProject#image}.
	Image *BedrockDataAutomationProjectOverrideConfigurationImage `field:"optional" json:"image" yaml:"image"`
	// Modality routing configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_automation_project#modality_routing BedrockDataAutomationProject#modality_routing}
	ModalityRouting *BedrockDataAutomationProjectOverrideConfigurationModalityRouting `field:"optional" json:"modalityRouting" yaml:"modalityRouting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_automation_project#video BedrockDataAutomationProject#video}.
	Video *BedrockDataAutomationProjectOverrideConfigurationVideo `field:"optional" json:"video" yaml:"video"`
}


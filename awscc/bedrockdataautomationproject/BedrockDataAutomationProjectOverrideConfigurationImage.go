// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectOverrideConfigurationImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_automation_project#modality_processing BedrockDataAutomationProject#modality_processing}.
	ModalityProcessing *BedrockDataAutomationProjectOverrideConfigurationImageModalityProcessing `field:"optional" json:"modalityProcessing" yaml:"modalityProcessing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_automation_project#sensitive_data_configuration BedrockDataAutomationProject#sensitive_data_configuration}.
	SensitiveDataConfiguration *BedrockDataAutomationProjectOverrideConfigurationImageSensitiveDataConfiguration `field:"optional" json:"sensitiveDataConfiguration" yaml:"sensitiveDataConfiguration"`
}


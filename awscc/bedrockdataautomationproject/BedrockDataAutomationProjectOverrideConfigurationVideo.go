// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectOverrideConfigurationVideo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_automation_project#modality_processing BedrockDataAutomationProject#modality_processing}.
	ModalityProcessing *BedrockDataAutomationProjectOverrideConfigurationVideoModalityProcessing `field:"optional" json:"modalityProcessing" yaml:"modalityProcessing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_automation_project#sensitive_data_configuration BedrockDataAutomationProject#sensitive_data_configuration}.
	SensitiveDataConfiguration *BedrockDataAutomationProjectOverrideConfigurationVideoSensitiveDataConfiguration `field:"optional" json:"sensitiveDataConfiguration" yaml:"sensitiveDataConfiguration"`
}


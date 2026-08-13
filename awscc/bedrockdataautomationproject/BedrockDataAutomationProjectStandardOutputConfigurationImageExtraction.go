// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationImageExtraction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#bounding_box BedrockDataAutomationProject#bounding_box}.
	BoundingBox *BedrockDataAutomationProjectStandardOutputConfigurationImageExtractionBoundingBox `field:"optional" json:"boundingBox" yaml:"boundingBox"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_automation_project#category BedrockDataAutomationProject#category}.
	Category *BedrockDataAutomationProjectStandardOutputConfigurationImageExtractionCategory `field:"optional" json:"category" yaml:"category"`
}


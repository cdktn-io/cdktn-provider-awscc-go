// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationVideoExtraction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#bounding_box BedrockDataAutomationProject#bounding_box}.
	BoundingBox *BedrockDataAutomationProjectStandardOutputConfigurationVideoExtractionBoundingBox `field:"optional" json:"boundingBox" yaml:"boundingBox"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_project#category BedrockDataAutomationProject#category}.
	Category *BedrockDataAutomationProjectStandardOutputConfigurationVideoExtractionCategory `field:"optional" json:"category" yaml:"category"`
}


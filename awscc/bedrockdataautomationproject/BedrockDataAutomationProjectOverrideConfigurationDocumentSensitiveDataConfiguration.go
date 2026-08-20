// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectOverrideConfigurationDocumentSensitiveDataConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_data_automation_project#detection_mode BedrockDataAutomationProject#detection_mode}.
	DetectionMode *string `field:"optional" json:"detectionMode" yaml:"detectionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_data_automation_project#detection_scope BedrockDataAutomationProject#detection_scope}.
	DetectionScope *[]*string `field:"optional" json:"detectionScope" yaml:"detectionScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_data_automation_project#pii_entities_configuration BedrockDataAutomationProject#pii_entities_configuration}.
	PiiEntitiesConfiguration *BedrockDataAutomationProjectOverrideConfigurationDocumentSensitiveDataConfigurationPiiEntitiesConfiguration `field:"optional" json:"piiEntitiesConfiguration" yaml:"piiEntitiesConfiguration"`
}


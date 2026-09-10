// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationPiiEntitiesConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_automation_project#pii_entity_types BedrockDataAutomationProject#pii_entity_types}.
	PiiEntityTypes *[]*string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_automation_project#redaction_mask_mode BedrockDataAutomationProject#redaction_mask_mode}.
	RedactionMaskMode *string `field:"optional" json:"redactionMaskMode" yaml:"redactionMaskMode"`
}


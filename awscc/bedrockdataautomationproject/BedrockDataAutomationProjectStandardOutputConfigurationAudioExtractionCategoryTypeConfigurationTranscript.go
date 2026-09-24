// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscript struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_automation_project#channel_labeling BedrockDataAutomationProject#channel_labeling}.
	ChannelLabeling *BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptChannelLabeling `field:"optional" json:"channelLabeling" yaml:"channelLabeling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_automation_project#speaker_labeling BedrockDataAutomationProject#speaker_labeling}.
	SpeakerLabeling *BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptSpeakerLabeling `field:"optional" json:"speakerLabeling" yaml:"speakerLabeling"`
}


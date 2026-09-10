// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdlpsetting


type QuicksightDlpSettingProviderConfigMicrosoftPurview struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_dlp_setting#credentials QuicksightDlpSetting#credentials}.
	Credentials *QuicksightDlpSettingProviderConfigMicrosoftPurviewCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_dlp_setting#label_action_mappings QuicksightDlpSetting#label_action_mappings}.
	LabelActionMappings interface{} `field:"optional" json:"labelActionMappings" yaml:"labelActionMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_dlp_setting#unmapped_action QuicksightDlpSetting#unmapped_action}.
	UnmappedAction *string `field:"optional" json:"unmappedAction" yaml:"unmappedAction"`
}


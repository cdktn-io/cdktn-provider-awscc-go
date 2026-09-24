// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdlpsetting


type QuicksightDlpSettingProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dlp_setting#microsoft_purview QuicksightDlpSetting#microsoft_purview}.
	MicrosoftPurview *QuicksightDlpSettingProviderConfigMicrosoftPurview `field:"optional" json:"microsoftPurview" yaml:"microsoftPurview"`
}


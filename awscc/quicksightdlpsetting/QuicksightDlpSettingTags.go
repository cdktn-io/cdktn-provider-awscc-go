// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdlpsetting


type QuicksightDlpSettingTags struct {
	// <p>Tag key.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_dlp_setting#key QuicksightDlpSetting#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// <p>Tag value.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_dlp_setting#value QuicksightDlpSetting#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


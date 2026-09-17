// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdlpsetting

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDlpSettingConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dlp_setting#dlp_setting_id QuicksightDlpSetting#dlp_setting_id}.
	DlpSettingId *string `field:"required" json:"dlpSettingId" yaml:"dlpSettingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dlp_setting#name QuicksightDlpSetting#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dlp_setting#provider_config QuicksightDlpSetting#provider_config}.
	ProviderConfig *QuicksightDlpSettingProviderConfig `field:"required" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dlp_setting#provider_outage_action QuicksightDlpSetting#provider_outage_action}.
	ProviderOutageAction *string `field:"required" json:"providerOutageAction" yaml:"providerOutageAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dlp_setting#provider_type QuicksightDlpSetting#provider_type}.
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dlp_setting#aws_account_id QuicksightDlpSetting#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dlp_setting#enabled QuicksightDlpSetting#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_dlp_setting#tags QuicksightDlpSetting#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


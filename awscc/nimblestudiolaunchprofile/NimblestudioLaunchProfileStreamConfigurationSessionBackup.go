// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfigurationSessionBackup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/nimblestudio_launch_profile#max_backups_to_retain NimblestudioLaunchProfile#max_backups_to_retain}.
	MaxBackupsToRetain *float64 `field:"optional" json:"maxBackupsToRetain" yaml:"maxBackupsToRetain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/nimblestudio_launch_profile#mode NimblestudioLaunchProfile#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}


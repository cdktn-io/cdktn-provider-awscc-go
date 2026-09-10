// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackApplicationSettings struct {
	// Enables or disables persistent application settings for users during their streaming sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appstream_stack#enabled AppstreamStack#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The path prefix for the S3 bucket where users' persistent application settings are stored.
	//
	// You can allow the same persistent application settings to be used across multiple stacks by specifying the same settings group for each stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appstream_stack#settings_group AppstreamStack#settings_group}
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
}


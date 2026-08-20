// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationset


type SesConfigurationSetVdmOptions struct {
	// Preferences regarding the Dashboard feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_configuration_set#dashboard_options SesConfigurationSet#dashboard_options}
	DashboardOptions *SesConfigurationSetVdmOptionsDashboardOptions `field:"optional" json:"dashboardOptions" yaml:"dashboardOptions"`
	// Preferences regarding the Guardian feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_configuration_set#guardian_options SesConfigurationSet#guardian_options}
	GuardianOptions *SesConfigurationSetVdmOptionsGuardianOptions `field:"optional" json:"guardianOptions" yaml:"guardianOptions"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2cisscanconfiguration


type Inspectorv2CisScanConfigurationScheduleWeeklyStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_cis_scan_configuration#time_of_day Inspectorv2CisScanConfiguration#time_of_day}.
	TimeOfDay *string `field:"optional" json:"timeOfDay" yaml:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/inspectorv2_cis_scan_configuration#time_zone Inspectorv2CisScanConfiguration#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}


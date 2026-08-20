// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2cisscanconfiguration


type Inspectorv2CisScanConfigurationScheduleMonthly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_cis_scan_configuration#day Inspectorv2CisScanConfiguration#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_cis_scan_configuration#start_time Inspectorv2CisScanConfiguration#start_time}.
	StartTime *Inspectorv2CisScanConfigurationScheduleMonthlyStartTime `field:"optional" json:"startTime" yaml:"startTime"`
}


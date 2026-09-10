// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicense


type LicensemanagerLicenseConsumptionConfigurationBorrowConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_license#allow_early_check_in LicensemanagerLicense#allow_early_check_in}.
	AllowEarlyCheckIn interface{} `field:"optional" json:"allowEarlyCheckIn" yaml:"allowEarlyCheckIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_license#max_time_to_live_in_minutes LicensemanagerLicense#max_time_to_live_in_minutes}.
	MaxTimeToLiveInMinutes *float64 `field:"optional" json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}


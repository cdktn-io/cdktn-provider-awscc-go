// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicense


type LicensemanagerLicenseEntitlements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/licensemanager_license#name LicensemanagerLicense#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/licensemanager_license#unit LicensemanagerLicense#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/licensemanager_license#allow_check_in LicensemanagerLicense#allow_check_in}.
	AllowCheckIn interface{} `field:"optional" json:"allowCheckIn" yaml:"allowCheckIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/licensemanager_license#max_count LicensemanagerLicense#max_count}.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/licensemanager_license#overage LicensemanagerLicense#overage}.
	Overage interface{} `field:"optional" json:"overage" yaml:"overage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/licensemanager_license#value LicensemanagerLicense#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicense


type LicensemanagerLicenseConsumptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/licensemanager_license#borrow_configuration LicensemanagerLicense#borrow_configuration}.
	BorrowConfiguration *LicensemanagerLicenseConsumptionConfigurationBorrowConfiguration `field:"optional" json:"borrowConfiguration" yaml:"borrowConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/licensemanager_license#provisional_configuration LicensemanagerLicense#provisional_configuration}.
	ProvisionalConfiguration *LicensemanagerLicenseConsumptionConfigurationProvisionalConfiguration `field:"optional" json:"provisionalConfiguration" yaml:"provisionalConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/licensemanager_license#renew_type LicensemanagerLicense#renew_type}.
	RenewType *string `field:"optional" json:"renewType" yaml:"renewType"`
}


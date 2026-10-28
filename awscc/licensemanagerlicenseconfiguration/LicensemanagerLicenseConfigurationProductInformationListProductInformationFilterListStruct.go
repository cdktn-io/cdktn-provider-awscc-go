// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseconfiguration


type LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStruct struct {
	// Logical operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#product_information_filter_comparator LicensemanagerLicenseConfiguration#product_information_filter_comparator}
	ProductInformationFilterComparator *string `field:"optional" json:"productInformationFilterComparator" yaml:"productInformationFilterComparator"`
	// Filter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#product_information_filter_name LicensemanagerLicenseConfiguration#product_information_filter_name}
	ProductInformationFilterName *string `field:"optional" json:"productInformationFilterName" yaml:"productInformationFilterName"`
	// Filter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#product_information_filter_value LicensemanagerLicenseConfiguration#product_information_filter_value}
	ProductInformationFilterValue *[]*string `field:"optional" json:"productInformationFilterValue" yaml:"productInformationFilterValue"`
}


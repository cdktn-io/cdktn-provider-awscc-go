// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseconfiguration


type LicensemanagerLicenseConfigurationProductInformationListStruct struct {
	// A product information filter list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#product_information_filter_list LicensemanagerLicenseConfiguration#product_information_filter_list}
	ProductInformationFilterList interface{} `field:"optional" json:"productInformationFilterList" yaml:"productInformationFilterList"`
	// Resource type. The possible values are SSM_MANAGED | RDS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#resource_type LicensemanagerLicenseConfiguration#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}


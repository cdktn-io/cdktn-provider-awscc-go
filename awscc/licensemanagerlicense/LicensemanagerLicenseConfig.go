// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicense

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerLicenseConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Beneficiary of the license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#beneficiary LicensemanagerLicense#beneficiary}
	Beneficiary *string `field:"required" json:"beneficiary" yaml:"beneficiary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#consumption_configuration LicensemanagerLicense#consumption_configuration}.
	ConsumptionConfiguration *LicensemanagerLicenseConsumptionConfiguration `field:"required" json:"consumptionConfiguration" yaml:"consumptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#entitlements LicensemanagerLicense#entitlements}.
	Entitlements interface{} `field:"required" json:"entitlements" yaml:"entitlements"`
	// Home region for the created license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#home_region LicensemanagerLicense#home_region}
	HomeRegion *string `field:"required" json:"homeRegion" yaml:"homeRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#issuer LicensemanagerLicense#issuer}.
	Issuer *LicensemanagerLicenseIssuer `field:"required" json:"issuer" yaml:"issuer"`
	// Name for the created license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#license_name LicensemanagerLicense#license_name}
	LicenseName *string `field:"required" json:"licenseName" yaml:"licenseName"`
	// Product name for the created license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#product_name LicensemanagerLicense#product_name}
	ProductName *string `field:"required" json:"productName" yaml:"productName"`
	// ProductSKU of the license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#product_sku LicensemanagerLicense#product_sku}
	ProductSku *string `field:"required" json:"productSku" yaml:"productSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#validity LicensemanagerLicense#validity}.
	Validity *LicensemanagerLicenseValidity `field:"required" json:"validity" yaml:"validity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#license_metadata LicensemanagerLicense#license_metadata}.
	LicenseMetadata interface{} `field:"optional" json:"licenseMetadata" yaml:"licenseMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#status LicensemanagerLicense#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license#tags LicensemanagerLicense#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


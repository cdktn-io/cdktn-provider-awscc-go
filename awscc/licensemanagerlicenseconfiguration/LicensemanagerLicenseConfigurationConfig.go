// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerLicenseConfigurationConfig struct {
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
	// Dimension used to track the license inventory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#license_counting_type LicensemanagerLicenseConfiguration#license_counting_type}
	LicenseCountingType *string `field:"required" json:"licenseCountingType" yaml:"licenseCountingType"`
	// Name of the license configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#name LicensemanagerLicenseConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the license configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#description LicensemanagerLicenseConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When true, disassociates a resource when software is uninstalled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#disassociate_when_not_found LicensemanagerLicenseConfiguration#disassociate_when_not_found}
	DisassociateWhenNotFound interface{} `field:"optional" json:"disassociateWhenNotFound" yaml:"disassociateWhenNotFound"`
	// Number of licenses managed by the license configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#license_count LicensemanagerLicenseConfiguration#license_count}
	LicenseCount *float64 `field:"optional" json:"licenseCount" yaml:"licenseCount"`
	// Indicates whether hard or soft license enforcement is used. Exceeding a hard limit blocks the launch of new instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#license_count_hard_limit LicensemanagerLicenseConfiguration#license_count_hard_limit}
	LicenseCountHardLimit interface{} `field:"optional" json:"licenseCountHardLimit" yaml:"licenseCountHardLimit"`
	// License configuration expiry as an epoch timestamp (seconds since 1970-01-01T00:00:00Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#license_expiry LicensemanagerLicenseConfiguration#license_expiry}
	LicenseExpiry *float64 `field:"optional" json:"licenseExpiry" yaml:"licenseExpiry"`
	// License rules. The syntax is #name=value (for example, #allowedTenancy=EC2-DedicatedHost).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#license_rules LicensemanagerLicenseConfiguration#license_rules}
	LicenseRules *[]*string `field:"optional" json:"licenseRules" yaml:"licenseRules"`
	// Product information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#product_information_list LicensemanagerLicenseConfiguration#product_information_list}
	ProductInformationList interface{} `field:"optional" json:"productInformationList" yaml:"productInformationList"`
	// Tags to add to the license configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanager_license_configuration#tags LicensemanagerLicenseConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


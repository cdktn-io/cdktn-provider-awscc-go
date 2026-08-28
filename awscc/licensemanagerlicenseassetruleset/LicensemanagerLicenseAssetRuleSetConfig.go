// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseassetruleset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerLicenseAssetRuleSetConfig struct {
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
	// License asset ruleset name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/licensemanager_license_asset_rule_set#name LicensemanagerLicenseAssetRuleSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// License asset rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/licensemanager_license_asset_rule_set#rules LicensemanagerLicenseAssetRuleSet#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// License asset ruleset description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/licensemanager_license_asset_rule_set#description LicensemanagerLicenseAssetRuleSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to add to the license asset ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/licensemanager_license_asset_rule_set#tags LicensemanagerLicenseAssetRuleSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


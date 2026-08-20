// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseassetruleset


type LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseConfigurationRuleStatementMatchingRuleStatement struct {
	// Constraint (e.g. Equals, Not_Equals).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/licensemanager_license_asset_rule_set#constraint LicensemanagerLicenseAssetRuleSet#constraint}
	Constraint *string `field:"optional" json:"constraint" yaml:"constraint"`
	// Key to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/licensemanager_license_asset_rule_set#key_to_match LicensemanagerLicenseAssetRuleSet#key_to_match}
	KeyToMatch *string `field:"optional" json:"keyToMatch" yaml:"keyToMatch"`
	// Values to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/licensemanager_license_asset_rule_set#value_to_match LicensemanagerLicenseAssetRuleSet#value_to_match}
	ValueToMatch *[]*string `field:"optional" json:"valueToMatch" yaml:"valueToMatch"`
}


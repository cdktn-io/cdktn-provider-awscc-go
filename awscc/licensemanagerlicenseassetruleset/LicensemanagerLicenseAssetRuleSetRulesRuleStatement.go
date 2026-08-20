// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseassetruleset


type LicensemanagerLicenseAssetRuleSetRulesRuleStatement struct {
	// Instance rule statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/licensemanager_license_asset_rule_set#instance_rule_statement LicensemanagerLicenseAssetRuleSet#instance_rule_statement}
	InstanceRuleStatement *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatement `field:"optional" json:"instanceRuleStatement" yaml:"instanceRuleStatement"`
	// License configuration rule statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/licensemanager_license_asset_rule_set#license_configuration_rule_statement LicensemanagerLicenseAssetRuleSet#license_configuration_rule_statement}
	LicenseConfigurationRuleStatement *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseConfigurationRuleStatement `field:"optional" json:"licenseConfigurationRuleStatement" yaml:"licenseConfigurationRuleStatement"`
	// License rule statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/licensemanager_license_asset_rule_set#license_rule_statement LicensemanagerLicenseAssetRuleSet#license_rule_statement}
	LicenseRuleStatement *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatement `field:"optional" json:"licenseRuleStatement" yaml:"licenseRuleStatement"`
}


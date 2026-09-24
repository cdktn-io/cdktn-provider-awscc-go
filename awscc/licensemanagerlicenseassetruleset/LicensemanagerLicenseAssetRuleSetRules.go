// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseassetruleset


type LicensemanagerLicenseAssetRuleSetRules struct {
	// Rule statement. Specify exactly one of InstanceRuleStatement, LicenseRuleStatement, or LicenseConfigurationRuleStatement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/licensemanager_license_asset_rule_set#rule_statement LicensemanagerLicenseAssetRuleSet#rule_statement}
	RuleStatement *LicensemanagerLicenseAssetRuleSetRulesRuleStatement `field:"required" json:"ruleStatement" yaml:"ruleStatement"`
}


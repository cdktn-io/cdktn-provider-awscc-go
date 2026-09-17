// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewruleset


type DatabrewRulesetRulesColumnSelectors struct {
	// The name of a column from a dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_ruleset#name DatabrewRuleset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A regular expression for selecting a column from a dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_ruleset#regex DatabrewRuleset#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}


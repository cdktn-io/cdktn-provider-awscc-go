// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewruleset


type DatabrewRulesetRulesSubstitutionMap struct {
	// Value or column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/databrew_ruleset#value DatabrewRuleset#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/databrew_ruleset#value_reference DatabrewRuleset#value_reference}
	ValueReference *string `field:"optional" json:"valueReference" yaml:"valueReference"`
}


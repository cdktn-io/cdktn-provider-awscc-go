// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewruleset


type DatabrewRulesetRules struct {
	// Expression with rule conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_ruleset#check_expression DatabrewRuleset#check_expression}
	CheckExpression *string `field:"required" json:"checkExpression" yaml:"checkExpression"`
	// Name of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_ruleset#name DatabrewRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_ruleset#column_selectors DatabrewRuleset#column_selectors}.
	ColumnSelectors interface{} `field:"optional" json:"columnSelectors" yaml:"columnSelectors"`
	// Boolean value to disable/enable a rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_ruleset#disabled DatabrewRuleset#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_ruleset#substitution_map DatabrewRuleset#substitution_map}.
	SubstitutionMap interface{} `field:"optional" json:"substitutionMap" yaml:"substitutionMap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/databrew_ruleset#threshold DatabrewRuleset#threshold}.
	Threshold *DatabrewRulesetRulesThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}


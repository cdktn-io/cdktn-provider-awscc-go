// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewruleset


type DatabrewRulesetRulesThreshold struct {
	// Threshold type for a rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_ruleset#type DatabrewRuleset#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Threshold unit for a rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_ruleset#unit DatabrewRuleset#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Threshold value for a rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_ruleset#value DatabrewRuleset#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


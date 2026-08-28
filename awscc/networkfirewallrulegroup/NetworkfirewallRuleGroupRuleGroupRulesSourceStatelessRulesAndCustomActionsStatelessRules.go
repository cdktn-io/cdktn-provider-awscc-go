// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#priority NetworkfirewallRuleGroup#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#rule_definition NetworkfirewallRuleGroup#rule_definition}.
	RuleDefinition *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinition `field:"optional" json:"ruleDefinition" yaml:"ruleDefinition"`
}


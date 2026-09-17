// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_rule_group#action NetworkfirewallRuleGroup#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_rule_group#header NetworkfirewallRuleGroup#header}.
	Header *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesHeader `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_rule_group#rule_options NetworkfirewallRuleGroup#rule_options}.
	RuleOptions interface{} `field:"optional" json:"ruleOptions" yaml:"ruleOptions"`
}


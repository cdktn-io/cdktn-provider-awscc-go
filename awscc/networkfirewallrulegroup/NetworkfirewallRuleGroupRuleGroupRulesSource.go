// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkfirewall_rule_group#rules_source_list NetworkfirewallRuleGroup#rules_source_list}.
	RulesSourceList *NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStruct `field:"optional" json:"rulesSourceList" yaml:"rulesSourceList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkfirewall_rule_group#rules_string NetworkfirewallRuleGroup#rules_string}.
	RulesString *string `field:"optional" json:"rulesString" yaml:"rulesString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkfirewall_rule_group#stateful_rules NetworkfirewallRuleGroup#stateful_rules}.
	StatefulRules interface{} `field:"optional" json:"statefulRules" yaml:"statefulRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkfirewall_rule_group#stateless_rules_and_custom_actions NetworkfirewallRuleGroup#stateless_rules_and_custom_actions}.
	StatelessRulesAndCustomActions *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions `field:"optional" json:"statelessRulesAndCustomActions" yaml:"statelessRulesAndCustomActions"`
}


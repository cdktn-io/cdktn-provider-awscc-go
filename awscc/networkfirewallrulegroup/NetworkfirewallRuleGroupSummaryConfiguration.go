// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupSummaryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkfirewall_rule_group#rule_options NetworkfirewallRuleGroup#rule_options}.
	RuleOptions *[]*string `field:"optional" json:"ruleOptions" yaml:"ruleOptions"`
}


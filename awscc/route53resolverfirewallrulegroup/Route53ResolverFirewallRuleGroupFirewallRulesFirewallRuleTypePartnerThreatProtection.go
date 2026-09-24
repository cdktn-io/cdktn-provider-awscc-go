// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverfirewallrulegroup


type Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypePartnerThreatProtection struct {
	// The partner identifier value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53resolver_firewall_rule_group#partner Route53ResolverFirewallRuleGroup#partner}
	Partner *string `field:"optional" json:"partner" yaml:"partner"`
}


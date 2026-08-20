// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverfirewallrulegroup


type Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleType struct {
	// Configuration for an advanced content category rule type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53resolver_firewall_rule_group#firewall_advanced_content_category Route53ResolverFirewallRuleGroup#firewall_advanced_content_category}
	FirewallAdvancedContentCategory *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedContentCategory `field:"optional" json:"firewallAdvancedContentCategory" yaml:"firewallAdvancedContentCategory"`
	// Configuration for an advanced threat category rule type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53resolver_firewall_rule_group#firewall_advanced_threat_category Route53ResolverFirewallRuleGroup#firewall_advanced_threat_category}
	FirewallAdvancedThreatCategory *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypeFirewallAdvancedThreatCategory `field:"optional" json:"firewallAdvancedThreatCategory" yaml:"firewallAdvancedThreatCategory"`
	// Configuration for a partner threat protection rule type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53resolver_firewall_rule_group#partner_threat_protection Route53ResolverFirewallRuleGroup#partner_threat_protection}
	PartnerThreatProtection *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleTypePartnerThreatProtection `field:"optional" json:"partnerThreatProtection" yaml:"partnerThreatProtection"`
}


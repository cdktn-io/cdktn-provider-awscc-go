// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverfirewallrulegroup


type Route53ResolverFirewallRuleGroupFirewallRules struct {
	// Rule Action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#action Route53ResolverFirewallRuleGroup#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// BlockOverrideDnsType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#block_override_dns_type Route53ResolverFirewallRuleGroup#block_override_dns_type}
	BlockOverrideDnsType *string `field:"optional" json:"blockOverrideDnsType" yaml:"blockOverrideDnsType"`
	// BlockOverrideDomain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#block_override_domain Route53ResolverFirewallRuleGroup#block_override_domain}
	BlockOverrideDomain *string `field:"optional" json:"blockOverrideDomain" yaml:"blockOverrideDomain"`
	// BlockOverrideTtl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#block_override_ttl Route53ResolverFirewallRuleGroup#block_override_ttl}
	BlockOverrideTtl *float64 `field:"optional" json:"blockOverrideTtl" yaml:"blockOverrideTtl"`
	// BlockResponse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#block_response Route53ResolverFirewallRuleGroup#block_response}
	BlockResponse *string `field:"optional" json:"blockResponse" yaml:"blockResponse"`
	// ConfidenceThreshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#confidence_threshold Route53ResolverFirewallRuleGroup#confidence_threshold}
	ConfidenceThreshold *string `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// DnsThreatProtection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#dns_threat_protection Route53ResolverFirewallRuleGroup#dns_threat_protection}
	DnsThreatProtection *string `field:"optional" json:"dnsThreatProtection" yaml:"dnsThreatProtection"`
	// ResourceId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#firewall_domain_list_id Route53ResolverFirewallRuleGroup#firewall_domain_list_id}
	FirewallDomainListId *string `field:"optional" json:"firewallDomainListId" yaml:"firewallDomainListId"`
	// FirewallDomainRedirectionAction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#firewall_domain_redirection_action Route53ResolverFirewallRuleGroup#firewall_domain_redirection_action}
	FirewallDomainRedirectionAction *string `field:"optional" json:"firewallDomainRedirectionAction" yaml:"firewallDomainRedirectionAction"`
	// Advanced firewall rule type. Mutually exclusive with FirewallDomainListId and DnsThreatProtection/ConfidenceThreshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#firewall_rule_type Route53ResolverFirewallRuleGroup#firewall_rule_type}
	FirewallRuleType *Route53ResolverFirewallRuleGroupFirewallRulesFirewallRuleType `field:"optional" json:"firewallRuleType" yaml:"firewallRuleType"`
	// ResourceId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#firewall_threat_protection_id Route53ResolverFirewallRuleGroup#firewall_threat_protection_id}
	FirewallThreatProtectionId *string `field:"optional" json:"firewallThreatProtectionId" yaml:"firewallThreatProtectionId"`
	// Rule Priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#priority Route53ResolverFirewallRuleGroup#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Qtype.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#qtype Route53ResolverFirewallRuleGroup#qtype}
	Qtype *string `field:"optional" json:"qtype" yaml:"qtype"`
	// The status of the firewall rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group#status Route53ResolverFirewallRuleGroup#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


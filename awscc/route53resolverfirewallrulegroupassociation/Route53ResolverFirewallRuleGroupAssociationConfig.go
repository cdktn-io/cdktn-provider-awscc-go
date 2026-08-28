// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverfirewallrulegroupassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53ResolverFirewallRuleGroupAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// FirewallRuleGroupId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group_association#firewall_rule_group_id Route53ResolverFirewallRuleGroupAssociation#firewall_rule_group_id}
	FirewallRuleGroupId *string `field:"required" json:"firewallRuleGroupId" yaml:"firewallRuleGroupId"`
	// Priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group_association#priority Route53ResolverFirewallRuleGroupAssociation#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// VpcId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group_association#vpc_id Route53ResolverFirewallRuleGroupAssociation#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// MutationProtectionStatus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group_association#mutation_protection Route53ResolverFirewallRuleGroupAssociation#mutation_protection}
	MutationProtection *string `field:"optional" json:"mutationProtection" yaml:"mutationProtection"`
	// FirewallRuleGroupAssociationName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group_association#name Route53ResolverFirewallRuleGroupAssociation#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_firewall_rule_group_association#tags Route53ResolverFirewallRuleGroupAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


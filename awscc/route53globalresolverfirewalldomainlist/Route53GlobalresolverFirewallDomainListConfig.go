// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53globalresolverfirewalldomainlist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53GlobalresolverFirewallDomainListConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53globalresolver_firewall_domain_list#global_resolver_id Route53GlobalresolverFirewallDomainList#global_resolver_id}.
	GlobalResolverId *string `field:"required" json:"globalResolverId" yaml:"globalResolverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53globalresolver_firewall_domain_list#name Route53GlobalresolverFirewallDomainList#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53globalresolver_firewall_domain_list#client_token Route53GlobalresolverFirewallDomainList#client_token}.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53globalresolver_firewall_domain_list#description Route53GlobalresolverFirewallDomainList#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// S3 URL to import domains from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53globalresolver_firewall_domain_list#domain_file_url Route53GlobalresolverFirewallDomainList#domain_file_url}
	DomainFileUrl *string `field:"optional" json:"domainFileUrl" yaml:"domainFileUrl"`
	// An inline list of domains to use for this domain list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53globalresolver_firewall_domain_list#domains Route53GlobalresolverFirewallDomainList#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53globalresolver_firewall_domain_list#tags Route53GlobalresolverFirewallDomainList#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


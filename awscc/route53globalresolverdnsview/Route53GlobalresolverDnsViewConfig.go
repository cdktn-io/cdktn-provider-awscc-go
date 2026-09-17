// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53globalresolverdnsview

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53GlobalresolverDnsViewConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53globalresolver_dns_view#global_resolver_id Route53GlobalresolverDnsView#global_resolver_id}.
	GlobalResolverId *string `field:"required" json:"globalResolverId" yaml:"globalResolverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53globalresolver_dns_view#name Route53GlobalresolverDnsView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53globalresolver_dns_view#client_token Route53GlobalresolverDnsView#client_token}.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53globalresolver_dns_view#description Route53GlobalresolverDnsView#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53globalresolver_dns_view#dnssec_validation Route53GlobalresolverDnsView#dnssec_validation}.
	DnssecValidation *string `field:"optional" json:"dnssecValidation" yaml:"dnssecValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53globalresolver_dns_view#edns_client_subnet Route53GlobalresolverDnsView#edns_client_subnet}.
	EdnsClientSubnet *string `field:"optional" json:"ednsClientSubnet" yaml:"ednsClientSubnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53globalresolver_dns_view#firewall_rules_fail_open Route53GlobalresolverDnsView#firewall_rules_fail_open}.
	FirewallRulesFailOpen *string `field:"optional" json:"firewallRulesFailOpen" yaml:"firewallRulesFailOpen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53globalresolver_dns_view#tags Route53GlobalresolverDnsView#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


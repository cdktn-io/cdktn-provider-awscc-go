// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverresolverrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53ResolverResolverRuleConfig struct {
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
	// When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD.
	//
	// When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_resolver_rule#rule_type Route53ResolverResolverRule#rule_type}
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// The name server domain for queries to be delegated to if a query matches the delegation record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_resolver_rule#delegation_record Route53ResolverResolverRule#delegation_record}
	DelegationRecord *string `field:"optional" json:"delegationRecord" yaml:"delegationRecord"`
	// DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_resolver_rule#domain_name Route53ResolverResolverRule#domain_name}
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The name for the Resolver rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_resolver_rule#name Route53ResolverResolverRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the endpoint that the rule is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_resolver_rule#resolver_endpoint_id Route53ResolverResolverRule#resolver_endpoint_id}
	ResolverEndpointId *string `field:"optional" json:"resolverEndpointId" yaml:"resolverEndpointId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_resolver_rule#tags Route53ResolverResolverRule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to.
	//
	// Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53resolver_resolver_rule#target_ips Route53ResolverResolverRule#target_ips}
	TargetIps interface{} `field:"optional" json:"targetIps" yaml:"targetIps"`
}


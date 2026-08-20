// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverresolverrule


type Route53ResolverResolverRuleTargetIps struct {
	// One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53resolver_resolver_rule#ip Route53ResolverResolverRule#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// One IPv6 address that you want to forward DNS queries to. You can specify only IPv6 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53resolver_resolver_rule#ipv_6 Route53ResolverResolverRule#ipv_6}
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// The port at Ip that you want to forward DNS queries to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53resolver_resolver_rule#port Route53ResolverResolverRule#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol that you want to use to forward DNS queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53resolver_resolver_rule#protocol Route53ResolverResolverRule#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The SNI of the target name servers for DoH/DoH-FIPS outbound endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53resolver_resolver_rule#server_name_indication Route53ResolverResolverRule#server_name_indication}
	ServerNameIndication *string `field:"optional" json:"serverNameIndication" yaml:"serverNameIndication"`
}


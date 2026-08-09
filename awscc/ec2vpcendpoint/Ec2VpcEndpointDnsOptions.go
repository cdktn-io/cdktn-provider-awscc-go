// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpcendpoint


type Ec2VpcEndpointDnsOptions struct {
	// The DNS records created for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_vpc_endpoint#dns_record_ip_type Ec2VpcEndpoint#dns_record_ip_type}
	DnsRecordIpType *string `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
	// Indicates whether to enable private DNS only for inbound endpoints.
	//
	// This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_vpc_endpoint#private_dns_only_for_inbound_resolver_endpoint Ec2VpcEndpoint#private_dns_only_for_inbound_resolver_endpoint}
	PrivateDnsOnlyForInboundResolverEndpoint *string `field:"optional" json:"privateDnsOnlyForInboundResolverEndpoint" yaml:"privateDnsOnlyForInboundResolverEndpoint"`
	// The preference for which private domains have a private hosted zone created for and associated with the specified VPC.
	//
	// Only supported when private DNS is enabled and when the VPC endpoint type is ServiceNetwork or Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_vpc_endpoint#private_dns_preference Ec2VpcEndpoint#private_dns_preference}
	PrivateDnsPreference *string `field:"optional" json:"privateDnsPreference" yaml:"privateDnsPreference"`
	// Indicates which of the private domains to create private hosted zones for and associate with the specified VPC.
	//
	// Only supported when private DNS is enabled and the private DNS preference is ``VERIFIED_DOMAINS_AND_SPECIFIED_DOMAINS`` or ``SPECIFIED_DOMAINS_ONLY``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_vpc_endpoint#private_dns_specified_domains Ec2VpcEndpoint#private_dns_specified_domains}
	PrivateDnsSpecifiedDomains *[]*string `field:"optional" json:"privateDnsSpecifiedDomains" yaml:"privateDnsSpecifiedDomains"`
}


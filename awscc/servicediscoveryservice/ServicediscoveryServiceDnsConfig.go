// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicediscoveryservice


type ServicediscoveryServiceDnsConfig struct {
	// A list of DNS records associated with the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicediscovery_service#dns_records ServicediscoveryService#dns_records}
	DnsRecords interface{} `field:"optional" json:"dnsRecords" yaml:"dnsRecords"`
	// The ID of the namespace for the DNS configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicediscovery_service#namespace_id ServicediscoveryService#namespace_id}
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The routing policy to use for DNS queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicediscovery_service#routing_policy ServicediscoveryService#routing_policy}
	RoutingPolicy *string `field:"optional" json:"routingPolicy" yaml:"routingPolicy"`
}


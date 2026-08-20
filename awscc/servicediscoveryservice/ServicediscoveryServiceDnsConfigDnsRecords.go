// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicediscoveryservice


type ServicediscoveryServiceDnsConfigDnsRecords struct {
	// The time-to-live (TTL) for the DNS record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicediscovery_service#ttl ServicediscoveryService#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// The DNS record type (e.g., A, AAAA, SRV).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicediscovery_service#type ServicediscoveryService#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


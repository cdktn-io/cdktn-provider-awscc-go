// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildomain


type LightsailDomainDomainEntries struct {
	// The ID of the domain recordset entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_domain#id LightsailDomain#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, specifies whether the domain entry is an alias used by the Lightsail load balancer, Lightsail container service, Lightsail content delivery network (CDN) distribution, or another AWS resource.
	//
	// You can include an alias (A type) record in your request, which points to the DNS name of a load balancer, container service, CDN distribution, or other AWS resource and routes traffic to that resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_domain#is_alias LightsailDomain#is_alias}
	IsAlias interface{} `field:"optional" json:"isAlias" yaml:"isAlias"`
	// The name of the domain entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_domain#name LightsailDomain#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target AWS name server (e.g., ns-111.awsdns-11.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_domain#target LightsailDomain#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The type of domain entry (e.g., A, CNAME, MX, NS, SOA, SRV, TXT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_domain#type LightsailDomain#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaymulticastdomain


type Ec2TransitGatewayMulticastDomainOptions struct {
	// Indicates whether to automatically cross-account subnet associations that are associated with the transit gateway multicast domain.
	//
	// Valid Values: enable | disable
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_transit_gateway_multicast_domain#auto_accept_shared_associations Ec2TransitGatewayMulticastDomain#auto_accept_shared_associations}
	AutoAcceptSharedAssociations *string `field:"optional" json:"autoAcceptSharedAssociations" yaml:"autoAcceptSharedAssociations"`
	// Indicates whether Internet Group Management Protocol (IGMP) version 2 is turned on for the transit gateway multicast domain.
	//
	// Valid Values: enable | disable
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_transit_gateway_multicast_domain#igmpv_2_support Ec2TransitGatewayMulticastDomain#igmpv_2_support}
	Igmpv2Support *string `field:"optional" json:"igmpv2Support" yaml:"igmpv2Support"`
	// Indicates whether support for statically configuring transit gateway multicast group sources is turned on. Valid Values: enable | disable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_transit_gateway_multicast_domain#static_sources_support Ec2TransitGatewayMulticastDomain#static_sources_support}
	StaticSourcesSupport *string `field:"optional" json:"staticSourcesSupport" yaml:"staticSourcesSupport"`
}


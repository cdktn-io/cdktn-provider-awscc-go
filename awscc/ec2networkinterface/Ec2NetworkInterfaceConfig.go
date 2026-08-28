// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinterface

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2NetworkInterfaceConfig struct {
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
	// The ID of the subnet to associate with the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#subnet_id Ec2NetworkInterface#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#connection_tracking_specification Ec2NetworkInterface#connection_tracking_specification}.
	ConnectionTrackingSpecification *Ec2NetworkInterfaceConnectionTrackingSpecification `field:"optional" json:"connectionTrackingSpecification" yaml:"connectionTrackingSpecification"`
	// A description for the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#description Ec2NetworkInterface#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If you have instances or ENIs that rely on the IPv6 address not changing, to avoid disrupting traffic to instances or ENIs, you can enable a primary IPv6 address.
	//
	// Enable this option to automatically assign an IPv6 associated with the ENI attached to your instance to be the primary IPv6 address. When you enable an IPv6 address to be a primary IPv6, you cannot disable it. Traffic will be routed to the primary IPv6 address until the instance is terminated or the ENI is detached. If you have multiple IPv6 addresses associated with an ENI and you enable a primary IPv6 address, the first IPv6 address associated with the ENI becomes the primary IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#enable_primary_ipv_6 Ec2NetworkInterface#enable_primary_ipv_6}
	EnablePrimaryIpv6 interface{} `field:"optional" json:"enablePrimaryIpv6" yaml:"enablePrimaryIpv6"`
	// A list of security group IDs associated with this network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#group_set Ec2NetworkInterface#group_set}
	GroupSet *[]*string `field:"optional" json:"groupSet" yaml:"groupSet"`
	// Indicates the type of network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#interface_type Ec2NetworkInterface#interface_type}
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// The number of IPv4 prefixes to assign to a network interface.
	//
	// When you specify a number of IPv4 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /28 prefixes. You can't specify a count of IPv4 prefixes if you've specified one of the following: specific IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#ipv_4_prefix_count Ec2NetworkInterface#ipv_4_prefix_count}
	Ipv4PrefixCount *float64 `field:"optional" json:"ipv4PrefixCount" yaml:"ipv4PrefixCount"`
	// Assigns a list of IPv4 prefixes to the network interface.
	//
	// If you want EC2 to automatically assign IPv4 prefixes, use the Ipv4PrefixCount property and do not specify this property. Presently, only /28 prefixes are supported. You can't specify IPv4 prefixes if you've specified one of the following: a count of IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#ipv_4_prefixes Ec2NetworkInterface#ipv_4_prefixes}
	Ipv4Prefixes interface{} `field:"optional" json:"ipv4Prefixes" yaml:"ipv4Prefixes"`
	// The number of IPv6 addresses to assign to a network interface.
	//
	// Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#ipv_6_address_count Ec2NetworkInterface#ipv_6_address_count}
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface.
	//
	// If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#ipv_6_addresses Ec2NetworkInterface#ipv_6_addresses}
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The number of IPv6 prefixes to assign to a network interface.
	//
	// When you specify a number of IPv6 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /80 prefixes. You can't specify a count of IPv6 prefixes if you've specified one of the following: specific IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#ipv_6_prefix_count Ec2NetworkInterface#ipv_6_prefix_count}
	Ipv6PrefixCount *float64 `field:"optional" json:"ipv6PrefixCount" yaml:"ipv6PrefixCount"`
	// Assigns a list of IPv6 prefixes to the network interface.
	//
	// If you want EC2 to automatically assign IPv6 prefixes, use the Ipv6PrefixCount property and do not specify this property. Presently, only /80 prefixes are supported. You can't specify IPv6 prefixes if you've specified one of the following: a count of IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#ipv_6_prefixes Ec2NetworkInterface#ipv_6_prefixes}
	Ipv6Prefixes interface{} `field:"optional" json:"ipv6Prefixes" yaml:"ipv6Prefixes"`
	// Assigns a single private IP address to the network interface, which is used as the primary private IP address.
	//
	// If you want to specify multiple private IP address, use the PrivateIpAddresses property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#private_ip_address Ec2NetworkInterface#private_ip_address}
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Assigns a list of private IP addresses to the network interface.
	//
	// You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#private_ip_addresses Ec2NetworkInterface#private_ip_addresses}
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// Public IP DNS hostname type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#public_ip_dns_hostname_type_specification Ec2NetworkInterface#public_ip_dns_hostname_type_specification}
	PublicIpDnsHostnameTypeSpecification *string `field:"optional" json:"publicIpDnsHostnameTypeSpecification" yaml:"publicIpDnsHostnameTypeSpecification"`
	// The number of secondary private IPv4 addresses to assign to a network interface.
	//
	// When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#secondary_private_ip_address_count Ec2NetworkInterface#secondary_private_ip_address_count}
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// Indicates whether traffic to or from the instance is validated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#source_dest_check Ec2NetworkInterface#source_dest_check}
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// An arbitrary set of tags (key-value pairs) for this network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_network_interface#tags Ec2NetworkInterface#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


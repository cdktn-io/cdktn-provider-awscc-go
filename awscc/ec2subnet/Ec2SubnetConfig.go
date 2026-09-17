// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2subnet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2SubnetConfig struct {
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
	// The ID of the VPC the subnet is in.
	//
	// If you update this property, you must also update the ``CidrBlock`` property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#vpc_id Ec2Subnet#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Indicates whether a network interface created in this subnet receives an IPv6 address.
	//
	// The default value is ``false``.
	//  If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#assign_ipv_6_address_on_creation Ec2Subnet#assign_ipv_6_address_on_creation}
	AssignIpv6AddressOnCreation interface{} `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// The Availability Zone of the subnet.  If you update this property, you must also update the ``CidrBlock`` property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#availability_zone Ec2Subnet#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AZ ID of the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#availability_zone_id Ec2Subnet#availability_zone_id}
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The IPv4 CIDR block assigned to the subnet.
	//
	// If you update this property, we create a new subnet, and then delete the existing one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#cidr_block Ec2Subnet#cidr_block}
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.
	//
	// You must first configure a NAT gateway in a public subnet (separate from the subnet containing the IPv6-only workloads). For example, the subnet containing the NAT gateway should have a ``0.0.0.0/0`` route pointing to the internet gateway. For more information, see [Configure DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-nat64-dns64.html#nat-gateway-nat64-dns64-walkthrough) in the *User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#enable_dns_64 Ec2Subnet#enable_dns_64}
	EnableDns64 interface{} `field:"optional" json:"enableDns64" yaml:"enableDns64"`
	// Indicates the device position for local network interfaces in this subnet.
	//
	// For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#enable_lni_at_device_index Ec2Subnet#enable_lni_at_device_index}
	EnableLniAtDeviceIndex *float64 `field:"optional" json:"enableLniAtDeviceIndex" yaml:"enableLniAtDeviceIndex"`
	// An IPv4 IPAM pool ID for the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#ipv_4_ipam_pool_id Ec2Subnet#ipv_4_ipam_pool_id}
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// An IPv4 netmask length for the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#ipv_4_netmask_length Ec2Subnet#ipv_4_netmask_length}
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// The IPv6 CIDR block.  If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#ipv_6_cidr_block Ec2Subnet#ipv_6_cidr_block}
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// An IPv6 IPAM pool ID for the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#ipv_6_ipam_pool_id Ec2Subnet#ipv_6_ipam_pool_id}
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#ipv_6_native Ec2Subnet#ipv_6_native}
	Ipv6Native interface{} `field:"optional" json:"ipv6Native" yaml:"ipv6Native"`
	// An IPv6 netmask length for the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#ipv_6_netmask_length Ec2Subnet#ipv_6_netmask_length}
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// Indicates whether instances launched in this subnet receive a public IPv4 address.
	//
	// The default value is ``false``.
	//  AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#map_public_ip_on_launch Ec2Subnet#map_public_ip_on_launch}
	MapPublicIpOnLaunch interface{} `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// The Amazon Resource Name (ARN) of the Outpost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#outpost_arn Ec2Subnet#outpost_arn}
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled.
	//
	// For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.
	//  Available options:
	//   +  EnableResourceNameDnsAAAARecord (true | false)
	//   +  EnableResourceNameDnsARecord (true | false)
	//   +  HostnameType (ip-name | resource-name)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#private_dns_name_options_on_launch Ec2Subnet#private_dns_name_options_on_launch}
	PrivateDnsNameOptionsOnLaunch *Ec2SubnetPrivateDnsNameOptionsOnLaunch `field:"optional" json:"privateDnsNameOptionsOnLaunch" yaml:"privateDnsNameOptionsOnLaunch"`
	// Any tags assigned to the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_subnet#tags Ec2Subnet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


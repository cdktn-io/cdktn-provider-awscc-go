// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2natgateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2NatGatewayConfig struct {
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
	// [Public NAT gateway only] The allocation ID of the Elastic IP address that's associated with the NAT gateway.
	//
	// This property is required for a public NAT gateway and cannot be specified with a private NAT gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#allocation_id Ec2NatGateway#allocation_id}
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Indicates whether this is a zonal (single-AZ) or regional (multi-AZ) NAT gateway.
	//
	// A zonal NAT gateway is a NAT Gateway that provides redundancy and scalability within a single availability zone. A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.
	//  For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the *Amazon VPC User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#availability_mode Ec2NatGateway#availability_mode}
	AvailabilityMode *string `field:"optional" json:"availabilityMode" yaml:"availabilityMode"`
	// For regional NAT gateways only: Specifies which Availability Zones you want the NAT gateway to support and the Elastic IP addresses (EIPs) to use in each AZ.
	//
	// The regional NAT gateway uses these EIPs to handle outbound NAT traffic from their respective AZs. If not specified, the NAT gateway will automatically expand to new AZs and associate EIPs upon detection of an elastic network interface. If you specify this parameter, auto-expansion is disabled and you must manually manage AZ coverage.
	//  A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.
	//  For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the *Amazon VPC User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#availability_zone_addresses Ec2NatGateway#availability_zone_addresses}
	AvailabilityZoneAddresses interface{} `field:"optional" json:"availabilityZoneAddresses" yaml:"availabilityZoneAddresses"`
	// Indicates whether the NAT gateway supports public or private connectivity. The default is public connectivity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#connectivity_type Ec2NatGateway#connectivity_type}
	ConnectivityType *string `field:"optional" json:"connectivityType" yaml:"connectivityType"`
	// The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress.
	//
	// Default value is 350 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#max_drain_duration_seconds Ec2NatGateway#max_drain_duration_seconds}
	MaxDrainDurationSeconds *float64 `field:"optional" json:"maxDrainDurationSeconds" yaml:"maxDrainDurationSeconds"`
	// The private IPv4 address to assign to the NAT gateway.
	//
	// If you don't provide an address, a private IPv4 address will be automatically assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#private_ip_address Ec2NatGateway#private_ip_address}
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Secondary EIP allocation IDs. For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html) in the *Amazon VPC User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#secondary_allocation_ids Ec2NatGateway#secondary_allocation_ids}
	SecondaryAllocationIds *[]*string `field:"optional" json:"secondaryAllocationIds" yaml:"secondaryAllocationIds"`
	// [Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway.
	//
	// For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
	//  ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#secondary_private_ip_address_count Ec2NatGateway#secondary_private_ip_address_count}
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// Secondary private IPv4 addresses.
	//
	// For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
	//  ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#secondary_private_ip_addresses Ec2NatGateway#secondary_private_ip_addresses}
	SecondaryPrivateIpAddresses *[]*string `field:"optional" json:"secondaryPrivateIpAddresses" yaml:"secondaryPrivateIpAddresses"`
	// The ID of the subnet in which the NAT gateway is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#subnet_id Ec2NatGateway#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags for the NAT gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#tags Ec2NatGateway#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC in which the NAT gateway is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_nat_gateway#vpc_id Ec2NatGateway#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}


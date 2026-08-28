// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2localgatewayroute

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2LocalGatewayRouteConfig struct {
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
	// The CIDR block used for destination matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_route#destination_cidr_block Ec2LocalGatewayRoute#destination_cidr_block}
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the local gateway route table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_route#local_gateway_route_table_id Ec2LocalGatewayRoute#local_gateway_route_table_id}
	LocalGatewayRouteTableId *string `field:"optional" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
	// The ID of the virtual interface group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_route#local_gateway_virtual_interface_group_id Ec2LocalGatewayRoute#local_gateway_virtual_interface_group_id}
	LocalGatewayVirtualInterfaceGroupId *string `field:"optional" json:"localGatewayVirtualInterfaceGroupId" yaml:"localGatewayVirtualInterfaceGroupId"`
	// The ID of the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_route#network_interface_id Ec2LocalGatewayRoute#network_interface_id}
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}


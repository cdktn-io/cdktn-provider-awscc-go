// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectprivatevirtualinterface

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectPrivateVirtualInterfaceConfig struct {
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
	// The BGP peers configured on this virtual interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#bgp_peers DirectconnectPrivateVirtualInterface#bgp_peers}
	BgpPeers interface{} `field:"required" json:"bgpPeers" yaml:"bgpPeers"`
	// The ID or ARN of the connection or LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#connection_id DirectconnectPrivateVirtualInterface#connection_id}
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// The name of the virtual interface assigned by the customer network.
	//
	// The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#virtual_interface_name DirectconnectPrivateVirtualInterface#virtual_interface_name}
	VirtualInterfaceName *string `field:"required" json:"virtualInterfaceName" yaml:"virtualInterfaceName"`
	// The ID of the VLAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#vlan DirectconnectPrivateVirtualInterface#vlan}
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// The Amazon Resource Name (ARN) of the role to allocate the private virtual interface.
	//
	// Needs directconnect:AllocatePrivateVirtualInterface permissions and tag permissions if applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#allocate_private_virtual_interface_role_arn DirectconnectPrivateVirtualInterface#allocate_private_virtual_interface_role_arn}
	AllocatePrivateVirtualInterfaceRoleArn *string `field:"optional" json:"allocatePrivateVirtualInterfaceRoleArn" yaml:"allocatePrivateVirtualInterfaceRoleArn"`
	// The ID or ARN of the Direct Connect gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#direct_connect_gateway_id DirectconnectPrivateVirtualInterface#direct_connect_gateway_id}
	DirectConnectGatewayId *string `field:"optional" json:"directConnectGatewayId" yaml:"directConnectGatewayId"`
	// Indicates whether to enable or disable SiteLink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#enable_site_link DirectconnectPrivateVirtualInterface#enable_site_link}
	EnableSiteLink interface{} `field:"optional" json:"enableSiteLink" yaml:"enableSiteLink"`
	// The maximum transmission unit (MTU), in bytes. The supported values are 1500 and 9001. The default value is 1500.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#mtu DirectconnectPrivateVirtualInterface#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// The tags associated with the private virtual interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#tags DirectconnectPrivateVirtualInterface#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ID or ARN of the virtual private gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/directconnect_private_virtual_interface#virtual_gateway_id DirectconnectPrivateVirtualInterface#virtual_gateway_id}
	VirtualGatewayId *string `field:"optional" json:"virtualGatewayId" yaml:"virtualGatewayId"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnecttransitvirtualinterface

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectTransitVirtualInterfaceConfig struct {
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
	// The BGP peers configured on this virtual interface..
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#bgp_peers DirectconnectTransitVirtualInterface#bgp_peers}
	BgpPeers interface{} `field:"required" json:"bgpPeers" yaml:"bgpPeers"`
	// The ID or ARN of the connection or LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#connection_id DirectconnectTransitVirtualInterface#connection_id}
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// The ID or ARN of the Direct Connect gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#direct_connect_gateway_id DirectconnectTransitVirtualInterface#direct_connect_gateway_id}
	DirectConnectGatewayId *string `field:"required" json:"directConnectGatewayId" yaml:"directConnectGatewayId"`
	// The name of the virtual interface assigned by the customer network.
	//
	// The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#virtual_interface_name DirectconnectTransitVirtualInterface#virtual_interface_name}
	VirtualInterfaceName *string `field:"required" json:"virtualInterfaceName" yaml:"virtualInterfaceName"`
	// The ID of the VLAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#vlan DirectconnectTransitVirtualInterface#vlan}
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// The Amazon Resource Name (ARN) of the role to allocate the TransitVifAllocation.
	//
	// Needs directconnect:AllocateTransitVirtualInterface permissions and tag permissions if applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#allocate_transit_virtual_interface_role_arn DirectconnectTransitVirtualInterface#allocate_transit_virtual_interface_role_arn}
	AllocateTransitVirtualInterfaceRoleArn *string `field:"optional" json:"allocateTransitVirtualInterfaceRoleArn" yaml:"allocateTransitVirtualInterfaceRoleArn"`
	// Indicates whether to enable or disable SiteLink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#enable_site_link DirectconnectTransitVirtualInterface#enable_site_link}
	EnableSiteLink interface{} `field:"optional" json:"enableSiteLink" yaml:"enableSiteLink"`
	// The maximum transmission unit (MTU), in bytes. The supported values are 1500 and 9001. The default value is 1500.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#mtu DirectconnectTransitVirtualInterface#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// The rate limit (bandwidth allocation) for the virtual interface.
	//
	// The value must be one of the supported bandwidth values (e.g., 50Mbps, 1Gbps, 10Gbps) and cannot exceed the bandwidth of the parent connection or LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#rate_limit DirectconnectTransitVirtualInterface#rate_limit}
	RateLimit *string `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// The tags associated with the private virtual interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/directconnect_transit_virtual_interface#tags DirectconnectTransitVirtualInterface#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


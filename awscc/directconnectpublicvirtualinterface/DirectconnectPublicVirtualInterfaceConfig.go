// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectpublicvirtualinterface

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectPublicVirtualInterfaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#bgp_peers DirectconnectPublicVirtualInterface#bgp_peers}
	BgpPeers interface{} `field:"required" json:"bgpPeers" yaml:"bgpPeers"`
	// The ID or ARN of the connection or LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#connection_id DirectconnectPublicVirtualInterface#connection_id}
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// The name of the virtual interface assigned by the customer network.
	//
	// The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#virtual_interface_name DirectconnectPublicVirtualInterface#virtual_interface_name}
	VirtualInterfaceName *string `field:"required" json:"virtualInterfaceName" yaml:"virtualInterfaceName"`
	// The ID of the VLAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#vlan DirectconnectPublicVirtualInterface#vlan}
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// The Amazon Resource Name (ARN) of the role to allocate the public virtual interface.
	//
	// Needs directconnect:AllocatePublicVirtualInterface permissions and tag permissions if applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#allocate_public_virtual_interface_role_arn DirectconnectPublicVirtualInterface#allocate_public_virtual_interface_role_arn}
	AllocatePublicVirtualInterfaceRoleArn *string `field:"optional" json:"allocatePublicVirtualInterfaceRoleArn" yaml:"allocatePublicVirtualInterfaceRoleArn"`
	// The rate limit (bandwidth allocation) for the virtual interface.
	//
	// The value must be one of the supported bandwidth values (e.g., 50Mbps, 1Gbps, 10Gbps) and cannot exceed the bandwidth of the parent connection or LAG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#rate_limit DirectconnectPublicVirtualInterface#rate_limit}
	RateLimit *string `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// The routes to be advertised to the AWS network in this region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#route_filter_prefixes DirectconnectPublicVirtualInterface#route_filter_prefixes}
	RouteFilterPrefixes *[]*string `field:"optional" json:"routeFilterPrefixes" yaml:"routeFilterPrefixes"`
	// The tags associated with the public virtual interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#tags DirectconnectPublicVirtualInterface#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbodbpeeringconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbOdbPeeringConnectionConfig struct {
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
	// The additional CIDR blocks for the ODB peering connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/odb_odb_peering_connection#additional_peer_network_cidrs OdbOdbPeeringConnection#additional_peer_network_cidrs}
	AdditionalPeerNetworkCidrs *[]*string `field:"optional" json:"additionalPeerNetworkCidrs" yaml:"additionalPeerNetworkCidrs"`
	// The name of the ODB peering connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/odb_odb_peering_connection#display_name OdbOdbPeeringConnection#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The unique identifier of the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/odb_odb_peering_connection#odb_network_id OdbOdbPeeringConnection#odb_network_id}
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// The unique identifier of the peer network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/odb_odb_peering_connection#peer_network_id OdbOdbPeeringConnection#peer_network_id}
	PeerNetworkId *string `field:"optional" json:"peerNetworkId" yaml:"peerNetworkId"`
	// The unique identifier of the VPC route table for which a route to the ODB network is automatically created during peering connection establishment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/odb_odb_peering_connection#peer_network_route_table_ids OdbOdbPeeringConnection#peer_network_route_table_ids}
	PeerNetworkRouteTableIds *[]*string `field:"optional" json:"peerNetworkRouteTableIds" yaml:"peerNetworkRouteTableIds"`
	// Tags to assign to the Odb peering connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/odb_odb_peering_connection#tags OdbOdbPeeringConnection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnectpeer

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerConnectPeerConfig struct {
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
	// The ID of the attachment to connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkmanager_connect_peer#connect_attachment_id NetworkmanagerConnectPeer#connect_attachment_id}
	ConnectAttachmentId *string `field:"required" json:"connectAttachmentId" yaml:"connectAttachmentId"`
	// The IP address of the Connect peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkmanager_connect_peer#peer_address NetworkmanagerConnectPeer#peer_address}
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// Bgp options for connect peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkmanager_connect_peer#bgp_options NetworkmanagerConnectPeer#bgp_options}
	BgpOptions *NetworkmanagerConnectPeerBgpOptions `field:"optional" json:"bgpOptions" yaml:"bgpOptions"`
	// The IP address of a core network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkmanager_connect_peer#core_network_address NetworkmanagerConnectPeer#core_network_address}
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// The inside IP addresses used for a Connect peer configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkmanager_connect_peer#inside_cidr_blocks NetworkmanagerConnectPeer#inside_cidr_blocks}
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// The subnet ARN for the connect peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkmanager_connect_peer#subnet_arn NetworkmanagerConnectPeer#subnet_arn}
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkmanager_connect_peer#tags NetworkmanagerConnectPeer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


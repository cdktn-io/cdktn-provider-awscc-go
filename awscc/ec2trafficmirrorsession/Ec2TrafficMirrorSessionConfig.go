// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2trafficmirrorsession

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2TrafficMirrorSessionConfig struct {
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
	// The ID of the source network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#network_interface_id Ec2TrafficMirrorSession#network_interface_id}
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions.
	//
	// The first session with a matching filter is the one that mirrors the packets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#session_number Ec2TrafficMirrorSession#session_number}
	SessionNumber *float64 `field:"required" json:"sessionNumber" yaml:"sessionNumber"`
	// The ID of a Traffic Mirror filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#traffic_mirror_filter_id Ec2TrafficMirrorSession#traffic_mirror_filter_id}
	TrafficMirrorFilterId *string `field:"required" json:"trafficMirrorFilterId" yaml:"trafficMirrorFilterId"`
	// The ID of a Traffic Mirror target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#traffic_mirror_target_id Ec2TrafficMirrorSession#traffic_mirror_target_id}
	TrafficMirrorTargetId *string `field:"required" json:"trafficMirrorTargetId" yaml:"trafficMirrorTargetId"`
	// The description of the Traffic Mirror session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#description Ec2TrafficMirrorSession#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the account that owns the Traffic Mirror session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#owner_id Ec2TrafficMirrorSession#owner_id}
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// The number of bytes in each packet to mirror.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#packet_length Ec2TrafficMirrorSession#packet_length}
	PacketLength *float64 `field:"optional" json:"packetLength" yaml:"packetLength"`
	// The tags assigned to the Traffic Mirror session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#tags Ec2TrafficMirrorSession#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The VXLAN ID for the Traffic Mirror session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_traffic_mirror_session#virtual_network_id Ec2TrafficMirrorSession#virtual_network_id}
	VirtualNetworkId *float64 `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}


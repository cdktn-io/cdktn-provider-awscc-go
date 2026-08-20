// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnectattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerConnectAttachmentConfig struct {
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
	// ID of the CoreNetwork that the attachment will be attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#core_network_id NetworkmanagerConnectAttachment#core_network_id}
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Edge location of the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#edge_location NetworkmanagerConnectAttachment#edge_location}
	EdgeLocation *string `field:"required" json:"edgeLocation" yaml:"edgeLocation"`
	// Protocol options for connect attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#options NetworkmanagerConnectAttachment#options}
	Options *NetworkmanagerConnectAttachmentOptions `field:"required" json:"options" yaml:"options"`
	// Id of transport attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#transport_attachment_id NetworkmanagerConnectAttachment#transport_attachment_id}
	TransportAttachmentId *string `field:"required" json:"transportAttachmentId" yaml:"transportAttachmentId"`
	// The name of the network function group attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#network_function_group_name NetworkmanagerConnectAttachment#network_function_group_name}
	NetworkFunctionGroupName *string `field:"optional" json:"networkFunctionGroupName" yaml:"networkFunctionGroupName"`
	// The attachment to move from one network function group to another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#proposed_network_function_group_change NetworkmanagerConnectAttachment#proposed_network_function_group_change}
	ProposedNetworkFunctionGroupChange *NetworkmanagerConnectAttachmentProposedNetworkFunctionGroupChange `field:"optional" json:"proposedNetworkFunctionGroupChange" yaml:"proposedNetworkFunctionGroupChange"`
	// The attachment to move from one segment to another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#proposed_segment_change NetworkmanagerConnectAttachment#proposed_segment_change}
	ProposedSegmentChange *NetworkmanagerConnectAttachmentProposedSegmentChange `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// Routing policy label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#routing_policy_label NetworkmanagerConnectAttachment#routing_policy_label}
	RoutingPolicyLabel *string `field:"optional" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
	// Tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_connect_attachment#tags NetworkmanagerConnectAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


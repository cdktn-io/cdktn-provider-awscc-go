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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#core_network_id NetworkmanagerConnectAttachment#core_network_id}
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Edge location of the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#edge_location NetworkmanagerConnectAttachment#edge_location}
	EdgeLocation *string `field:"required" json:"edgeLocation" yaml:"edgeLocation"`
	// Protocol options for connect attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#options NetworkmanagerConnectAttachment#options}
	Options *NetworkmanagerConnectAttachmentOptions `field:"required" json:"options" yaml:"options"`
	// Id of transport attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#transport_attachment_id NetworkmanagerConnectAttachment#transport_attachment_id}
	TransportAttachmentId *string `field:"required" json:"transportAttachmentId" yaml:"transportAttachmentId"`
	// The complete set of NAT gateway rules on the attachment.
	//
	// Updating this property replaces the entire rule set, and rules left out of the new value are removed. To remove all rules, specify an empty list or remove this property. NAT gateway support can't be enabled without rules, so it is then disabled unless NatGatewaySupport is specified. Only the core network owner can configure NAT gateway rules. The rules aren't returned when the attachment is read, so drift detection doesn't check them. Not available in AWS GovCloud (US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#nat_gateway_rules NetworkmanagerConnectAttachment#nat_gateway_rules}
	NatGatewayRules interface{} `field:"optional" json:"natGatewayRules" yaml:"natGatewayRules"`
	// Whether NAT gateway support is enabled on the attachment.
	//
	// Required when NatGatewayRules lists rules. Set to DISABLE with NatGatewayRules to keep the rules on the attachment without translating traffic, and set it to ENABLE to activate them again. An attachment that is pending acceptance applies this setting once it is accepted. Only the core network owner can configure or read NAT gateway settings. Not available in AWS GovCloud (US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#nat_gateway_support NetworkmanagerConnectAttachment#nat_gateway_support}
	NatGatewaySupport *string `field:"optional" json:"natGatewaySupport" yaml:"natGatewaySupport"`
	// The name of the network function group attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#network_function_group_name NetworkmanagerConnectAttachment#network_function_group_name}
	NetworkFunctionGroupName *string `field:"optional" json:"networkFunctionGroupName" yaml:"networkFunctionGroupName"`
	// Routing policy label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#routing_policy_label NetworkmanagerConnectAttachment#routing_policy_label}
	RoutingPolicyLabel *string `field:"optional" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
	// Tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_connect_attachment#tags NetworkmanagerConnectAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


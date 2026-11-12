// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerdirectconnectgatewayattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerDirectConnectGatewayAttachmentConfig struct {
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
	// The ID of a core network for the Direct Connect Gateway attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#core_network_id NetworkmanagerDirectConnectGatewayAttachment#core_network_id}
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The ARN of the Direct Connect Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#direct_connect_gateway_arn NetworkmanagerDirectConnectGatewayAttachment#direct_connect_gateway_arn}
	DirectConnectGatewayArn *string `field:"required" json:"directConnectGatewayArn" yaml:"directConnectGatewayArn"`
	// The Regions where the edges are located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#edge_locations NetworkmanagerDirectConnectGatewayAttachment#edge_locations}
	EdgeLocations *[]*string `field:"required" json:"edgeLocations" yaml:"edgeLocations"`
	// The complete set of NAT gateway rules on the attachment.
	//
	// Updating this property replaces the entire rule set, and rules left out of the new value are removed. To remove all rules, specify an empty list or remove this property. NAT gateway support can't be enabled without rules, so it is then disabled unless NatGatewaySupport is specified. Only the core network owner can configure NAT gateway rules. The rules aren't returned when the attachment is read, so drift detection doesn't check them. Not available in AWS GovCloud (US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#nat_gateway_rules NetworkmanagerDirectConnectGatewayAttachment#nat_gateway_rules}
	NatGatewayRules interface{} `field:"optional" json:"natGatewayRules" yaml:"natGatewayRules"`
	// Whether NAT gateway support is enabled on the attachment.
	//
	// Required when NatGatewayRules lists rules. Set to DISABLE with NatGatewayRules to keep the rules on the attachment without translating traffic, and set it to ENABLE to activate them again. An attachment that is pending acceptance applies this setting once it is accepted. Only the core network owner can configure or read NAT gateway settings. Not available in AWS GovCloud (US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#nat_gateway_support NetworkmanagerDirectConnectGatewayAttachment#nat_gateway_support}
	NatGatewaySupport *string `field:"optional" json:"natGatewaySupport" yaml:"natGatewaySupport"`
	// Whether system routes are installed for the NAT CIDRs of the attachment's NAT gateway rules.
	//
	// If omitted, a new NAT gateway configuration uses the service default of true and an existing configuration keeps its current value, so removing this property from a template doesn't reset the setting. Only the core network owner can configure this setting. Not available in AWS GovCloud (US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#nat_system_route_support NetworkmanagerDirectConnectGatewayAttachment#nat_system_route_support}
	NatSystemRouteSupport interface{} `field:"optional" json:"natSystemRouteSupport" yaml:"natSystemRouteSupport"`
	// Routing policy label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#routing_policy_label NetworkmanagerDirectConnectGatewayAttachment#routing_policy_label}
	RoutingPolicyLabel *string `field:"optional" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
	// Tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#tags NetworkmanagerDirectConnectGatewayAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


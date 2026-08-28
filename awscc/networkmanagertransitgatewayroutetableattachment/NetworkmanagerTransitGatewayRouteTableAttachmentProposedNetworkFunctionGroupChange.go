// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagertransitgatewayroutetableattachment


type NetworkmanagerTransitGatewayRouteTableAttachmentProposedNetworkFunctionGroupChange struct {
	// The rule number in the policy document that applies to this change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#attachment_policy_rule_number NetworkmanagerTransitGatewayRouteTableAttachment#attachment_policy_rule_number}
	AttachmentPolicyRuleNumber *float64 `field:"optional" json:"attachmentPolicyRuleNumber" yaml:"attachmentPolicyRuleNumber"`
	// The name of the network function group to change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#network_function_group_name NetworkmanagerTransitGatewayRouteTableAttachment#network_function_group_name}
	NetworkFunctionGroupName *string `field:"optional" json:"networkFunctionGroupName" yaml:"networkFunctionGroupName"`
	// The key-value tags that changed for the network function group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#tags NetworkmanagerTransitGatewayRouteTableAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


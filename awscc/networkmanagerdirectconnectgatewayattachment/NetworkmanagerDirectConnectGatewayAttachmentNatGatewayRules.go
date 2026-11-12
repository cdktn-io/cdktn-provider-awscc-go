// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerdirectconnectgatewayattachment


type NetworkmanagerDirectConnectGatewayAttachmentNatGatewayRules struct {
	// The IPv4 CIDR that source traffic is matched against for translation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#match_source_cidr NetworkmanagerDirectConnectGatewayAttachment#match_source_cidr}
	MatchSourceCidr *string `field:"optional" json:"matchSourceCidr" yaml:"matchSourceCidr"`
	// The IPv4 CIDR that matched source traffic is translated into. Must have the same prefix length as MatchSourceCidr.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#nat_cidr NetworkmanagerDirectConnectGatewayAttachment#nat_cidr}
	NatCidr *string `field:"optional" json:"natCidr" yaml:"natCidr"`
	// The rule number that uniquely identifies the NAT rule on the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_direct_connect_gateway_attachment#rule_number NetworkmanagerDirectConnectGatewayAttachment#rule_number}
	RuleNumber *float64 `field:"optional" json:"ruleNumber" yaml:"ruleNumber"`
}


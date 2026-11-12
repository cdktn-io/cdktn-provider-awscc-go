// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagervpcattachment


type NetworkmanagerVpcAttachmentNatGatewayRules struct {
	// The IPv4 CIDR that source traffic is matched against for translation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_vpc_attachment#match_source_cidr NetworkmanagerVpcAttachment#match_source_cidr}
	MatchSourceCidr *string `field:"optional" json:"matchSourceCidr" yaml:"matchSourceCidr"`
	// The IPv4 CIDR that matched source traffic is translated into. Must have the same prefix length as MatchSourceCidr.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_vpc_attachment#nat_cidr NetworkmanagerVpcAttachment#nat_cidr}
	NatCidr *string `field:"optional" json:"natCidr" yaml:"natCidr"`
	// The rule number that uniquely identifies the NAT rule on the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_vpc_attachment#rule_number NetworkmanagerVpcAttachment#rule_number}
	RuleNumber *float64 `field:"optional" json:"ruleNumber" yaml:"ruleNumber"`
}


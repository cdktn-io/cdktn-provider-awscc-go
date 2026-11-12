// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewayattachmentnatconfiguration


type Ec2TransitGatewayAttachmentNatConfigurationNatGatewayRules struct {
	// The source CIDR that the rule matches on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_transit_gateway_attachment_nat_configuration#match_source_cidr Ec2TransitGatewayAttachmentNatConfiguration#match_source_cidr}
	MatchSourceCidr *string `field:"required" json:"matchSourceCidr" yaml:"matchSourceCidr"`
	// The CIDR that the matched source addresses are translated to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_transit_gateway_attachment_nat_configuration#nat_cidr Ec2TransitGatewayAttachmentNatConfiguration#nat_cidr}
	NatCidr *string `field:"required" json:"natCidr" yaml:"natCidr"`
	// The rule number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_transit_gateway_attachment_nat_configuration#rule_number Ec2TransitGatewayAttachmentNatConfiguration#rule_number}
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaypolicytableentry


type Ec2TransitGatewayPolicyTableEntryPolicyRule struct {
	// The destination CIDR block for the transit gateway policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#destination_cidr_block Ec2TransitGatewayPolicyTableEntry#destination_cidr_block}
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The destination port range for the transit gateway policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#destination_port_range Ec2TransitGatewayPolicyTableEntry#destination_port_range}
	DestinationPortRange *string `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// The protocol for the transit gateway policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#protocol Ec2TransitGatewayPolicyTableEntry#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The source CIDR block for the transit gateway policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#source_cidr_block Ec2TransitGatewayPolicyTableEntry#source_cidr_block}
	SourceCidrBlock *string `field:"optional" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// The source port range for the transit gateway policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_policy_table_entry#source_port_range Ec2TransitGatewayPolicyTableEntry#source_port_range}
	SourcePortRange *string `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}


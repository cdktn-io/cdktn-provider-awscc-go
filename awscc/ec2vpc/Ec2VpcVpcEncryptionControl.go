// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpc


type Ec2VpcVpcEncryptionControl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#egress_only_internet_gateway_exclusion Ec2Vpc#egress_only_internet_gateway_exclusion}.
	EgressOnlyInternetGatewayExclusion *string `field:"optional" json:"egressOnlyInternetGatewayExclusion" yaml:"egressOnlyInternetGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#elastic_file_system_exclusion Ec2Vpc#elastic_file_system_exclusion}.
	ElasticFileSystemExclusion *string `field:"optional" json:"elasticFileSystemExclusion" yaml:"elasticFileSystemExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#internet_gateway_exclusion Ec2Vpc#internet_gateway_exclusion}.
	InternetGatewayExclusion *string `field:"optional" json:"internetGatewayExclusion" yaml:"internetGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#lambda_exclusion Ec2Vpc#lambda_exclusion}.
	LambdaExclusion *string `field:"optional" json:"lambdaExclusion" yaml:"lambdaExclusion"`
	// The encryption mode for the VPC Encryption Control configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#mode Ec2Vpc#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#nat_gateway_exclusion Ec2Vpc#nat_gateway_exclusion}.
	NatGatewayExclusion *string `field:"optional" json:"natGatewayExclusion" yaml:"natGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#virtual_private_gateway_exclusion Ec2Vpc#virtual_private_gateway_exclusion}.
	VirtualPrivateGatewayExclusion *string `field:"optional" json:"virtualPrivateGatewayExclusion" yaml:"virtualPrivateGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#vpc_lattice_exclusion Ec2Vpc#vpc_lattice_exclusion}.
	VpcLatticeExclusion *string `field:"optional" json:"vpcLatticeExclusion" yaml:"vpcLatticeExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc#vpc_peering_exclusion Ec2Vpc#vpc_peering_exclusion}.
	VpcPeeringExclusion *string `field:"optional" json:"vpcPeeringExclusion" yaml:"vpcPeeringExclusion"`
}


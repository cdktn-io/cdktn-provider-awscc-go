// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpcencryptioncontrol

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VpcEncryptionControlConfig struct {
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
	// Used to enable or disable EIGW exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#egress_only_internet_gateway_exclusion_input Ec2VpcEncryptionControl#egress_only_internet_gateway_exclusion_input}
	EgressOnlyInternetGatewayExclusionInput *string `field:"optional" json:"egressOnlyInternetGatewayExclusionInput" yaml:"egressOnlyInternetGatewayExclusionInput"`
	// Used to enable or disable EFS exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#elastic_file_system_exclusion_input Ec2VpcEncryptionControl#elastic_file_system_exclusion_input}
	ElasticFileSystemExclusionInput *string `field:"optional" json:"elasticFileSystemExclusionInput" yaml:"elasticFileSystemExclusionInput"`
	// Used to enable or disable IGW exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#internet_gateway_exclusion_input Ec2VpcEncryptionControl#internet_gateway_exclusion_input}
	InternetGatewayExclusionInput *string `field:"optional" json:"internetGatewayExclusionInput" yaml:"internetGatewayExclusionInput"`
	// Used to enable or disable Lambda exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#lambda_exclusion_input Ec2VpcEncryptionControl#lambda_exclusion_input}
	LambdaExclusionInput *string `field:"optional" json:"lambdaExclusionInput" yaml:"lambdaExclusionInput"`
	// The VPC encryption control mode, either monitor or enforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#mode Ec2VpcEncryptionControl#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Used to enable or disable Nat gateway exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#nat_gateway_exclusion_input Ec2VpcEncryptionControl#nat_gateway_exclusion_input}
	NatGatewayExclusionInput *string `field:"optional" json:"natGatewayExclusionInput" yaml:"natGatewayExclusionInput"`
	// The tags to assign to the VPC encryption control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#tags Ec2VpcEncryptionControl#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Used to enable or disable VGW exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#virtual_private_gateway_exclusion_input Ec2VpcEncryptionControl#virtual_private_gateway_exclusion_input}
	VirtualPrivateGatewayExclusionInput *string `field:"optional" json:"virtualPrivateGatewayExclusionInput" yaml:"virtualPrivateGatewayExclusionInput"`
	// The VPC on which this VPC encryption control is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#vpc_id Ec2VpcEncryptionControl#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Used to enable or disable Vpc Lattice exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#vpc_lattice_exclusion_input Ec2VpcEncryptionControl#vpc_lattice_exclusion_input}
	VpcLatticeExclusionInput *string `field:"optional" json:"vpcLatticeExclusionInput" yaml:"vpcLatticeExclusionInput"`
	// Used to enable or disable VPC peering exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpc_encryption_control#vpc_peering_exclusion_input Ec2VpcEncryptionControl#vpc_peering_exclusion_input}
	VpcPeeringExclusionInput *string `field:"optional" json:"vpcPeeringExclusionInput" yaml:"vpcPeeringExclusionInput"`
}


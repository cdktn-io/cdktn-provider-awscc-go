// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce


type SagemakerWorkforceWorkforceVpcConfig struct {
	// The VPC security group IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_workforce#security_group_ids SagemakerWorkforce#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The VPC subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_workforce#subnets SagemakerWorkforce#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The ID of the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_workforce#vpc_id SagemakerWorkforce#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}


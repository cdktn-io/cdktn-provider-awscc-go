// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelVpcConfig struct {
	// The VPC security group IDs, in the form `sg-xxxxxxxx`.
	//
	// Specify the security groups for the VPC that is specified in the `Subnets` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#security_group_ids SagemakerModel#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect your training job or model.
	//
	// For information about the availability of specific instance types, see [Supported Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#subnets SagemakerModel#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}


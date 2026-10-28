// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceVpcLatticeConfigurations struct {
	// The advanced settings for VPC Lattice used in blue/green deployments.
	//
	// Specify the alternate target group and listener rules required for traffic shifting during blue/green deployments. For more information, see [Required resources for Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/blue-green-deployment-implementation.html) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#advanced_configuration EcsService#advanced_configuration}
	AdvancedConfiguration *EcsServiceVpcLatticeConfigurationsAdvancedConfiguration `field:"optional" json:"advancedConfiguration" yaml:"advancedConfiguration"`
	// The name of the port mapping to register in the VPC Lattice target group.
	//
	// This is the name of the ``portMapping`` you defined in your task definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#port_name EcsService#port_name}
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
	// The ARN of the IAM role to associate with this VPC Lattice configuration.
	//
	// This is the Amazon ECS infrastructure IAM role that is used to manage your VPC Lattice infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#role_arn EcsService#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The full Amazon Resource Name (ARN) of the target group or groups associated with the VPC Lattice configuration that the Amazon ECS tasks will be registered to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#target_group_arn EcsService#target_group_arn}
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}


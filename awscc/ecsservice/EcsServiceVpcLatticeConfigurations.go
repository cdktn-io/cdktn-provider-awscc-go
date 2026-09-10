// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceVpcLatticeConfigurations struct {
	// The name of the port mapping to register in the VPC Lattice target group.
	//
	// This is the name of the ``portMapping`` you defined in your task definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#port_name EcsService#port_name}
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
	// The ARN of the IAM role to associate with this VPC Lattice configuration.
	//
	// This is the Amazon ECS infrastructure IAM role that is used to manage your VPC Lattice infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#role_arn EcsService#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The full Amazon Resource Name (ARN) of the target group or groups associated with the VPC Lattice configuration that the Amazon ECS tasks will be registered to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#target_group_arn EcsService#target_group_arn}
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceNetworkConfigurationAwsvpcConfiguration struct {
	// Whether the task's elastic network interface receives a public IP address.
	//
	// Consider the following when you set this value:
	//   +  When you use ``create-service`` or ``update-service``, the default is ``DISABLED``.
	//   +  When the service ``deploymentController`` is ``ECS``, the value must be ``DISABLED``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#assign_public_ip EcsService#assign_public_ip}
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The IDs of the security groups associated with the task or service.
	//
	// If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified.
	//   All specified security groups must be from the same VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#security_groups EcsService#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the subnets associated with the task or service.
	//
	// There's a limit of 16 subnets that can be specified.
	//   All specified subnets must be from the same VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#subnets EcsService#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}


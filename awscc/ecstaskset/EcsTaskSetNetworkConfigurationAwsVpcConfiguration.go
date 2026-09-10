// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskset


type EcsTaskSetNetworkConfigurationAwsVpcConfiguration struct {
	// Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_set#assign_public_ip EcsTaskSet#assign_public_ip}
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The security groups associated with the task or service.
	//
	// If you do not specify a security group, the default security group for the VPC is used. There is a limit of 5 security groups that can be specified per AwsVpcConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_set#security_groups EcsTaskSet#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets associated with the task or service.
	//
	// There is a limit of 16 subnets that can be specified per AwsVpcConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_set#subnets EcsTaskSet#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceNetworkConfiguration struct {
	// The VPC subnets and security groups that are associated with a task.
	//
	// All specified subnets and security groups must be from the same VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_service#awsvpc_configuration EcsService#awsvpc_configuration}
	AwsvpcConfiguration *EcsServiceNetworkConfigurationAwsvpcConfiguration `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
}


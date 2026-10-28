// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceVpcLatticeConfigurationsAdvancedConfiguration struct {
	// The Amazon Resource Name (ARN) of the alternate target group associated with the VPC Lattice Configuration for Amazon ECS blue/green deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#alternate_target_group_arn EcsService#alternate_target_group_arn}
	AlternateTargetGroupArn *string `field:"optional" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// The Amazon Resource Name (ARN) that identifies the production listener rule or listener for routing production traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#production_listener_rule EcsService#production_listener_rule}
	ProductionListenerRule *string `field:"optional" json:"productionListenerRule" yaml:"productionListenerRule"`
	// The Amazon Resource Name (ARN) that identifies the test listener rule or listener for routing test traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#test_listener_rule EcsService#test_listener_rule}
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}


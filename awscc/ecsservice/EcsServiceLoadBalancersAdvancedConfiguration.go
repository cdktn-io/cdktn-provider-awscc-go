// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceLoadBalancersAdvancedConfiguration struct {
	// The Amazon Resource Name (ARN) of the alternate target group for Amazon ECS blue/green deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_service#alternate_target_group_arn EcsService#alternate_target_group_arn}
	AlternateTargetGroupArn *string `field:"optional" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// The Amazon Resource Name (ARN) that that identifies the production listener rule (in the case of an Application Load Balancer) or listener (in the case for an Network Load Balancer) for routing production traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_service#production_listener_rule EcsService#production_listener_rule}
	ProductionListenerRule *string `field:"optional" json:"productionListenerRule" yaml:"productionListenerRule"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon ECS permission to call the Elastic Load Balancing APIs for you.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_service#role_arn EcsService#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) that identifies ) that identifies the test listener rule (in the case of an Application Load Balancer) or listener (in the case for an Network Load Balancer) for routing test traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_service#test_listener_rule EcsService#test_listener_rule}
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}


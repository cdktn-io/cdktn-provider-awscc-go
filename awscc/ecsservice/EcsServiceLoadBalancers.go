// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceLoadBalancers struct {
	// The advanced settings for the load balancer used in blue/green deployments.
	//
	// Specify the alternate target group, listener rules, and IAM role required for traffic shifting during blue/green deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#advanced_configuration EcsService#advanced_configuration}
	AdvancedConfiguration *EcsServiceLoadBalancersAdvancedConfiguration `field:"optional" json:"advancedConfiguration" yaml:"advancedConfiguration"`
	// The name of the container (as it appears in a container definition) to associate with the load balancer.
	//
	// You need to specify the container name when configuring the target group for an Amazon ECS load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#container_name EcsService#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port on the container to associate with the load balancer.
	//
	// This port must correspond to a ``containerPort`` in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they're launched on must allow ingress traffic on the ``hostPort`` of the port mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#container_port EcsService#container_port}
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The name of the load balancer to associate with the Amazon ECS service or task set.
	//
	// If you are using an Application Load Balancer or a Network Load Balancer the load balancer name parameter should be omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#load_balancer_name EcsService#load_balancer_name}
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set.
	//
	// A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer.
	//  For services using the ``ECS`` deployment controller, you can specify one or multiple target groups. For more information, see [Registering multiple target groups with a service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html) in the *Amazon Elastic Container Service Developer Guide*.
	//  For services using the ``CODE_DEPLOY`` deployment controller, you're required to define two target groups for the load balancer. For more information, see [Blue/green deployment with CodeDeploy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html) in the *Amazon Elastic Container Service Developer Guide*.
	//   If your service's task definition uses the ``awsvpc`` network mode, you must choose ``ip`` as the target type, not ``instance``. Do this when creating your target groups because tasks that use the ``awsvpc`` network mode are associated with an elastic network interface, not an Amazon EC2 instance. This network mode is required for the Fargate launch type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_service#target_group_arn EcsService#target_group_arn}
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}


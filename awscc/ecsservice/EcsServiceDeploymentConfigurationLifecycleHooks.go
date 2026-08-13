// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationLifecycleHooks struct {
	// Use this field to specify custom parameters that ECS passes to your hook target invocations (such as a Lambda function).
	//
	// This field must be a JSON object as a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#hook_details EcsService#hook_details}
	HookDetails *string `field:"optional" json:"hookDetails" yaml:"hookDetails"`
	// The Amazon Resource Name (ARN) of the hook target.
	//
	// For ``AWS_LAMBDA`` hooks, this is the Lambda function ARN. This field is not applicable for ``PAUSE`` hooks.
	//  You must provide this parameter when configuring an ``AWS_LAMBDA`` lifecycle hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#hook_target_arn EcsService#hook_target_arn}
	HookTargetArn *string `field:"optional" json:"hookTargetArn" yaml:"hookTargetArn"`
	// The lifecycle stages at which to run the hook.
	//
	// Choose from these valid values:
	//   +  RECONCILE_SERVICE
	//  The reconciliation stage that only happens when you start a new service deployment with more than 1 service revision in an ACTIVE state.
	//  You can use a lifecycle hook for this stage.
	//   +  PRE_SCALE_UP
	//  The green service revision has not started. The blue service revision is handling 100% of the production traffic. There is no test traffic.
	//  You can use a lifecycle hook for this stage.
	//   +  POST_SCALE_UP
	//  The green service revision has started. The blue service revision is handling 100% of the production traffic. There is no test traffic.
	//  You can use a lifecycle hook for this stage.
	//   +  TEST_TRAFFIC_SHIFT
	//  The blue and green service revisions are running. The blue service revision handles 100% of the production traffic. The green service revision is migrating from 0% to 100% of test traffic.
	//  You can use a lifecycle hook for this stage.
	//   +  POST_TEST_TRAFFIC_SHIFT
	//  The test traffic shift is complete. The green service revision handles 100% of the test traffic.
	//  You can use a lifecycle hook for this stage.
	//   +  PRE_PRODUCTION_TRAFFIC_SHIFT
	//  Occurs before production traffic shift. For linear and canary deployments, this stage is invoked before every traffic shift step.
	//  You can use a lifecycle hook for this stage.
	//   +  PRODUCTION_TRAFFIC_SHIFT
	//  Production traffic is shifting to the green service revision. The green service revision is migrating from 0% to 100% of production traffic. For linear and canary deployments, this stage is invoked at every traffic shift step.
	//  You can use a lifecycle hook for this stage.
	//   +  POST_PRODUCTION_TRAFFIC_SHIFT
	//  The production traffic shift is complete.
	//  You can use a lifecycle hook for this stage.
	//
	//   ``PAUSE`` hooks cannot be configured at ``TEST_TRAFFIC_SHIFT`` or ``PRODUCTION_TRAFFIC_SHIFT`` stages. These stages are only valid for ``AWS_LAMBDA`` hooks.
	//   You must provide this parameter when configuring a deployment lifecycle hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#lifecycle_stages EcsService#lifecycle_stages}
	LifecycleStages *[]*string `field:"optional" json:"lifecycleStages" yaml:"lifecycleStages"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon ECS permission to call Lambda functions on your behalf.
	//
	// For more information, see [Permissions required for Lambda functions in Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/blue-green-permissions.html) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#role_arn EcsService#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The type of action the lifecycle hook performs.
	//
	// Valid values are:
	//   +  ``AWS_LAMBDA`` - Invokes a Lambda function at the specified lifecycle stage. This is the default value.
	//   +  ``PAUSE`` - Pauses the deployment at the specified lifecycle stage until you call ``ContinueServiceDeployment`` to continue or roll back.
	//
	//  This field is optional. If not specified, the default value is ``AWS_LAMBDA``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#target_type EcsService#target_type}
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// The timeout configuration for the lifecycle hook.
	//
	// This specifies how long Amazon ECS waits before taking the timeout action if the hook is not resolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#timeout_configuration EcsService#timeout_configuration}
	TimeoutConfiguration *EcsServiceDeploymentConfigurationLifecycleHooksTimeoutConfiguration `field:"optional" json:"timeoutConfiguration" yaml:"timeoutConfiguration"`
}


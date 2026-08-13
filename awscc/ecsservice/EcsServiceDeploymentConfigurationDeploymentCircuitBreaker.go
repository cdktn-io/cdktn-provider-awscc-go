// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationDeploymentCircuitBreaker struct {
	// Determines whether to use the deployment circuit breaker logic for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#enable EcsService#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Specifies whether the deployment circuit breaker resets its failure count when a task reaches a healthy state.
	//
	// When set to ``true``, a task that reaches a healthy state resets the failure count to ``0``. When set to ``false``, Amazon ECS does not reset the failure count. The default is ``true``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#reset_on_healthy_task EcsService#reset_on_healthy_task}
	ResetOnHealthyTask interface{} `field:"optional" json:"resetOnHealthyTask" yaml:"resetOnHealthyTask"`
	// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails.
	//
	// If rollback is on, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#rollback EcsService#rollback}
	Rollback interface{} `field:"optional" json:"rollback" yaml:"rollback"`
	// The threshold configuration that controls when the deployment circuit breaker triggers.
	//
	// The ``type`` and ``value`` together determine how many task failures are tolerated before the circuit breaker activates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#threshold_configuration EcsService#threshold_configuration}
	ThresholdConfiguration *EcsServiceDeploymentConfigurationDeploymentCircuitBreakerThresholdConfiguration `field:"optional" json:"thresholdConfiguration" yaml:"thresholdConfiguration"`
}


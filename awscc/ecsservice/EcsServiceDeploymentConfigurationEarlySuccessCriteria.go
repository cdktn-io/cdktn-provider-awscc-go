// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationEarlySuccessCriteria struct {
	// Specifies whether to use the early success criteria for the service deployment.
	//
	// When set to ``false``, the deployment uses the default behavior, where Amazon ECS considers the deployment successful when the target service revision fully stabilizes and the previous tasks are removed. The default value is ``false``.
	//  When set to ``true``, Amazon ECS monitors the deployment to meet early success criteria. You must also specify ``healthyPercent`` and ``sourceServiceRevisionCleanup``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#enable EcsService#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// The percentage of healthy tasks that the target service revision must reach before Amazon ECS considers the deployment successful.
	//
	// This percentage is relative to the service's ``desiredCount`` and must be an integer between ``0`` and ``100``. This value must be greater than or equal to the ``minimumHealthyPercent`` value.
	//  After this percentage of tasks is healthy and the bake time elapses, Amazon ECS completes the deployment. Amazon ECS continues to scale the target service revision to 100 percent in the background.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#healthy_percent EcsService#healthy_percent}
	HealthyPercent *float64 `field:"optional" json:"healthyPercent" yaml:"healthyPercent"`
	// The time when Amazon ECS removes the source revisions' tasks relative to deployment completion.
	//
	// The valid values are:
	//   +  ``BLOCKING``—Amazon ECS removes the previous tasks before it marks the deployment as successful.
	//   +  ``DEFERRED``—Amazon ECS marks the deployment successful, and then removes the previous tasks in the background.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#source_service_revision_cleanup EcsService#source_service_revision_cleanup}
	SourceServiceRevisionCleanup *string `field:"optional" json:"sourceServiceRevisionCleanup" yaml:"sourceServiceRevisionCleanup"`
}


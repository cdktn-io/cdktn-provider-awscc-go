// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationDeploymentCircuitBreakerThresholdConfiguration struct {
	// Determines how Amazon ECS uses ``value`` to calculate the failure threshold.
	//
	// For the percentage types (``BOUNDED_PERCENT`` and ``UNBOUNDED_PERCENT``), Amazon ECS multiplies ``value`` by the latest service desired count. For ``COUNT``, Amazon ECS uses ``value`` directly as the threshold. The default is ``BOUNDED_PERCENT``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_service#type EcsService#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the integer that Amazon ECS uses to calculate the failure threshold.
	//
	// When ``type`` is ``COUNT``, this value is the failure threshold itself. When ``type`` is a percentage type, Amazon ECS multiplies this value by the latest service desired count to produce the failure threshold. The default is ``50``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_service#value EcsService#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


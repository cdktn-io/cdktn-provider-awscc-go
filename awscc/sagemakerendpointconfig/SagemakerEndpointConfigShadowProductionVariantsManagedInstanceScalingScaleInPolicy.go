// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigShadowProductionVariantsManagedInstanceScalingScaleInPolicy struct {
	// The cooldown period, in minutes, after the last endpoint operation before the endpoint evaluates consolidation scale-in opportunities.
	//
	// Valid values are 5 to 1440. The default is 20.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint_config#cooldown_in_minutes SagemakerEndpointConfigA#cooldown_in_minutes}
	CooldownInMinutes *float64 `field:"optional" json:"cooldownInMinutes" yaml:"cooldownInMinutes"`
	// The maximum number of instances that the endpoint can terminate at a time during a consolidation scale-in operation.
	//
	// Valid values are 1 to 100. The default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint_config#maximum_step_size SagemakerEndpointConfigA#maximum_step_size}
	MaximumStepSize *float64 `field:"optional" json:"maximumStepSize" yaml:"maximumStepSize"`
	// The strategy for scaling in instances.
	//
	// IDLE_RELEASE releases instances that have no hosted inference component copies. CONSOLIDATION consolidates inference component copies onto fewer instances to release more instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint_config#strategy SagemakerEndpointConfigA#strategy}
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}


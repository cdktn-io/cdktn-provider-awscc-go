// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalabletarget


type ApplicationautoscalingScalableTargetSuspendedState struct {
	// Whether scale in by a target tracking scaling policy or a step scaling policy is suspended.
	//
	// Set the value to ``true`` if you don't want Application Auto Scaling to remove capacity when a scaling policy is triggered. The default is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationautoscaling_scalable_target#dynamic_scaling_in_suspended ApplicationautoscalingScalableTarget#dynamic_scaling_in_suspended}
	DynamicScalingInSuspended interface{} `field:"optional" json:"dynamicScalingInSuspended" yaml:"dynamicScalingInSuspended"`
	// Whether scale out by a target tracking scaling policy or a step scaling policy is suspended.
	//
	// Set the value to ``true`` if you don't want Application Auto Scaling to add capacity when a scaling policy is triggered. The default is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationautoscaling_scalable_target#dynamic_scaling_out_suspended ApplicationautoscalingScalableTarget#dynamic_scaling_out_suspended}
	DynamicScalingOutSuspended interface{} `field:"optional" json:"dynamicScalingOutSuspended" yaml:"dynamicScalingOutSuspended"`
	// Whether scheduled scaling is suspended.
	//
	// Set the value to ``true`` if you don't want Application Auto Scaling to add or remove capacity by initiating scheduled actions. The default is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationautoscaling_scalable_target#scheduled_scaling_suspended ApplicationautoscalingScalableTarget#scheduled_scaling_suspended}
	ScheduledScalingSuspended interface{} `field:"optional" json:"scheduledScalingSuspended" yaml:"scheduledScalingSuspended"`
}


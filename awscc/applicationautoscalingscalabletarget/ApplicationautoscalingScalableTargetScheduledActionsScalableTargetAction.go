// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalabletarget


type ApplicationautoscalingScalableTargetScheduledActionsScalableTargetAction struct {
	// The maximum capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationautoscaling_scalable_target#max_capacity ApplicationautoscalingScalableTarget#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationautoscaling_scalable_target#min_capacity ApplicationautoscalingScalableTarget#min_capacity}
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}


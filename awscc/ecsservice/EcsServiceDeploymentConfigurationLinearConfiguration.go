// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationLinearConfiguration struct {
	// The amount of time in minutes to wait between each traffic shifting step during a linear deployment.
	//
	// Valid values are 0 to 1440 minutes (24 hours). The default value is 6. This bake time is not applied after reaching 100 percent traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_service#step_bake_time_in_minutes EcsService#step_bake_time_in_minutes}
	StepBakeTimeInMinutes *float64 `field:"optional" json:"stepBakeTimeInMinutes" yaml:"stepBakeTimeInMinutes"`
	// The percentage of production traffic to shift in each step during a linear deployment.
	//
	// Valid values are multiples of 0.1 from 3.0 to 100.0. The default value is 10.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_service#step_percent EcsService#step_percent}
	StepPercent *float64 `field:"optional" json:"stepPercent" yaml:"stepPercent"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationCanaryConfiguration struct {
	// The amount of time in minutes to wait during the canary phase before shifting the remaining production traffic to the new service revision.
	//
	// Valid values are 0 to 1440 minutes (24 hours). The default value is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#canary_bake_time_in_minutes EcsService#canary_bake_time_in_minutes}
	CanaryBakeTimeInMinutes *float64 `field:"optional" json:"canaryBakeTimeInMinutes" yaml:"canaryBakeTimeInMinutes"`
	// The percentage of production traffic to shift to the new service revision during the canary phase.
	//
	// Valid values are multiples of 0.1 from 0.1 to 100.0. The default value is 5.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#canary_percent EcsService#canary_percent}
	CanaryPercent *float64 `field:"optional" json:"canaryPercent" yaml:"canaryPercent"`
}


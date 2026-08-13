// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet


type GameliftContainerFleetDeploymentConfiguration struct {
	// The strategy to apply in case of impairment; defaults to MAINTAIN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_container_fleet#impairment_strategy GameliftContainerFleet#impairment_strategy}
	ImpairmentStrategy *string `field:"optional" json:"impairmentStrategy" yaml:"impairmentStrategy"`
	// The minimum percentage of healthy required; defaults to 75.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_container_fleet#minimum_healthy_percentage GameliftContainerFleet#minimum_healthy_percentage}
	MinimumHealthyPercentage *float64 `field:"optional" json:"minimumHealthyPercentage" yaml:"minimumHealthyPercentage"`
	// The protection strategy for deployment on the container fleet; defaults to WITH_PROTECTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_container_fleet#protection_strategy GameliftContainerFleet#protection_strategy}
	ProtectionStrategy *string `field:"optional" json:"protectionStrategy" yaml:"protectionStrategy"`
}


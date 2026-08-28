// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet


type GameliftContainerFleetGameSessionCreationLimitPolicy struct {
	// The maximum number of game sessions that an individual can create during the policy period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_fleet#new_game_sessions_per_creator GameliftContainerFleet#new_game_sessions_per_creator}
	NewGameSessionsPerCreator *float64 `field:"optional" json:"newGameSessionsPerCreator" yaml:"newGameSessionsPerCreator"`
	// The time span used in evaluating the resource creation limit policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_fleet#policy_period_in_minutes GameliftContainerFleet#policy_period_in_minutes}
	PolicyPeriodInMinutes *float64 `field:"optional" json:"policyPeriodInMinutes" yaml:"policyPeriodInMinutes"`
}


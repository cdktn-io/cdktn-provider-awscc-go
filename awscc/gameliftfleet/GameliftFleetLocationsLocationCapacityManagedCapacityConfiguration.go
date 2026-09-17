// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetLocationsLocationCapacityManagedCapacityConfiguration struct {
	// Length of time, in minutes, that Amazon GameLift Servers will wait before scaling in your MinSize and DesiredInstances to 0 after a period with no game session activity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#scale_in_after_inactivity_minutes GameliftFleet#scale_in_after_inactivity_minutes}
	ScaleInAfterInactivityMinutes *float64 `field:"optional" json:"scaleInAfterInactivityMinutes" yaml:"scaleInAfterInactivityMinutes"`
	// The strategy Amazon GameLift Servers will use to automatically scale your capacity to and from zero in response to game session activity.
	//
	// Game session activity refers to any active running sessions or game session requests. When set to SCALE_TO_AND_FROM_ZERO, MinSize must not be specified and will be managed automatically. When set to MANUAL, MinSize is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#zero_capacity_strategy GameliftFleet#zero_capacity_strategy}
	ZeroCapacityStrategy *string `field:"optional" json:"zeroCapacityStrategy" yaml:"zeroCapacityStrategy"`
}


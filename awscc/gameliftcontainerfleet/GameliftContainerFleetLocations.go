// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet


type GameliftContainerFleetLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#location GameliftContainerFleet#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Current resource capacity settings in a specified fleet or location.
	//
	// The location value might refer to a fleet's remote location or its home Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#location_capacity GameliftContainerFleet#location_capacity}
	LocationCapacity *GameliftContainerFleetLocationsLocationCapacity `field:"optional" json:"locationCapacity" yaml:"locationCapacity"`
	// The player gateway status for the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#player_gateway_status GameliftContainerFleet#player_gateway_status}
	PlayerGatewayStatus *string `field:"optional" json:"playerGatewayStatus" yaml:"playerGatewayStatus"`
	// A list of fleet actions that have been suspended in the fleet location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#stopped_actions GameliftContainerFleet#stopped_actions}
	StoppedActions *[]*string `field:"optional" json:"stoppedActions" yaml:"stoppedActions"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#location GameliftFleet#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Current resource capacity settings in a specified fleet or location.
	//
	// The location value might refer to a fleet's remote location or its home Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#location_capacity GameliftFleet#location_capacity}
	LocationCapacity *GameliftFleetLocationsLocationCapacity `field:"optional" json:"locationCapacity" yaml:"locationCapacity"`
	// The player gateway status for the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#player_gateway_status GameliftFleet#player_gateway_status}
	PlayerGatewayStatus *string `field:"optional" json:"playerGatewayStatus" yaml:"playerGatewayStatus"`
}


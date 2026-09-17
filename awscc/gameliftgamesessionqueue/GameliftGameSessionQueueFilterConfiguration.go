// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgamesessionqueue


type GameliftGameSessionQueueFilterConfiguration struct {
	// A list of locations to allow game session placement in, in the form of AWS Region codes such as us-west-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_session_queue#allowed_locations GameliftGameSessionQueue#allowed_locations}
	AllowedLocations *[]*string `field:"optional" json:"allowedLocations" yaml:"allowedLocations"`
}


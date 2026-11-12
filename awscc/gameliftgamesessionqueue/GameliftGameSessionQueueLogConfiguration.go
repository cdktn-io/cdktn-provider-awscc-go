// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgamesessionqueue


type GameliftGameSessionQueueLogConfiguration struct {
	// The set of destinations that receive the game session queue's placement logs.
	//
	// A queue can deliver logs to between 1 and 3 destinations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/gamelift_game_session_queue#delivery_destinations GameliftGameSessionQueue#delivery_destinations}
	DeliveryDestinations interface{} `field:"optional" json:"deliveryDestinations" yaml:"deliveryDestinations"`
}


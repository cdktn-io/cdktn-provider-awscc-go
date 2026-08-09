// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetRuntimeConfiguration struct {
	// The maximum amount of time (in seconds) that a game session can remain in status ACTIVATING.
	//
	// If the game session is not active before the timeout, activation is terminated and the game session status is changed to TERMINATED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_fleet#game_session_activation_timeout_seconds GameliftFleet#game_session_activation_timeout_seconds}
	GameSessionActivationTimeoutSeconds *float64 `field:"optional" json:"gameSessionActivationTimeoutSeconds" yaml:"gameSessionActivationTimeoutSeconds"`
	// The maximum number of game sessions with status ACTIVATING to allow on an instance simultaneously.
	//
	// This setting limits the amount of instance resources that can be used for new game activations at any one time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_fleet#max_concurrent_game_session_activations GameliftFleet#max_concurrent_game_session_activations}
	MaxConcurrentGameSessionActivations *float64 `field:"optional" json:"maxConcurrentGameSessionActivations" yaml:"maxConcurrentGameSessionActivations"`
	// A collection of server process configurations that describe which server processes to run on each instance in a fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_fleet#server_processes GameliftFleet#server_processes}
	ServerProcesses interface{} `field:"optional" json:"serverProcesses" yaml:"serverProcesses"`
}


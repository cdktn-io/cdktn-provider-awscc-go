// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgameservergroup


type GameliftGameServerGroupAutoScalingPolicy struct {
	// Length of time, in seconds, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_game_server_group#estimated_instance_warmup GameliftGameServerGroup#estimated_instance_warmup}
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// Settings for a target-based scaling policy applied to Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_game_server_group#target_tracking_configuration GameliftGameServerGroup#target_tracking_configuration}
	TargetTrackingConfiguration *GameliftGameServerGroupAutoScalingPolicyTargetTrackingConfiguration `field:"optional" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgameservergroup


type GameliftGameServerGroupInstanceDefinitions struct {
	// An EC2 instance type designation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_game_server_group#instance_type GameliftGameServerGroup#instance_type}
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_game_server_group#weighted_capacity GameliftGameServerGroup#weighted_capacity}
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


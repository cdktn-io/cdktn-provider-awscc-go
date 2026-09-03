// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgameservergroup


type GameliftGameServerGroupTags struct {
	// The key for a developer-defined key:value pair for tagging an AWS resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_game_server_group#key GameliftGameServerGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for a developer-defined key:value pair for tagging an AWS resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_game_server_group#value GameliftGameServerGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


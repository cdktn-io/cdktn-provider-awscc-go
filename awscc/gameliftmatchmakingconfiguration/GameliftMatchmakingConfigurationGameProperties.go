// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftmatchmakingconfiguration


type GameliftMatchmakingConfigurationGameProperties struct {
	// The game property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_matchmaking_configuration#key GameliftMatchmakingConfiguration#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The game property value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_matchmaking_configuration#value GameliftMatchmakingConfiguration#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


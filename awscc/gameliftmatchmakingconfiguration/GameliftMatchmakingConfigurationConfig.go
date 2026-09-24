// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftmatchmakingconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftMatchmakingConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A flag that indicates whether a match that was created with this configuration must be accepted by the matched players.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#acceptance_required GameliftMatchmakingConfiguration#acceptance_required}
	AcceptanceRequired interface{} `field:"required" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// A unique identifier for the matchmaking configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#name GameliftMatchmakingConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#request_timeout_seconds GameliftMatchmakingConfiguration#request_timeout_seconds}
	RequestTimeoutSeconds *float64 `field:"required" json:"requestTimeoutSeconds" yaml:"requestTimeoutSeconds"`
	// A unique identifier for the matchmaking rule set to use with this configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#rule_set_name GameliftMatchmakingConfiguration#rule_set_name}
	RuleSetName *string `field:"required" json:"ruleSetName" yaml:"ruleSetName"`
	// The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#acceptance_timeout_seconds GameliftMatchmakingConfiguration#acceptance_timeout_seconds}
	AcceptanceTimeoutSeconds *float64 `field:"optional" json:"acceptanceTimeoutSeconds" yaml:"acceptanceTimeoutSeconds"`
	// The number of player slots in a match to keep open for future players.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#additional_player_count GameliftMatchmakingConfiguration#additional_player_count}
	AdditionalPlayerCount *float64 `field:"optional" json:"additionalPlayerCount" yaml:"additionalPlayerCount"`
	// The method used to backfill game sessions created with this matchmaking configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#backfill_mode GameliftMatchmakingConfiguration#backfill_mode}
	BackfillMode *string `field:"optional" json:"backfillMode" yaml:"backfillMode"`
	// A time stamp indicating when this data object was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#creation_time GameliftMatchmakingConfiguration#creation_time}
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Information to attach to all events related to the matchmaking configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#custom_event_data GameliftMatchmakingConfiguration#custom_event_data}
	CustomEventData *string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// A descriptive label that is associated with matchmaking configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#description GameliftMatchmakingConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#flex_match_mode GameliftMatchmakingConfiguration#flex_match_mode}
	FlexMatchMode *string `field:"optional" json:"flexMatchMode" yaml:"flexMatchMode"`
	// A set of custom properties for a game session, formatted as key:value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#game_properties GameliftMatchmakingConfiguration#game_properties}
	GameProperties interface{} `field:"optional" json:"gameProperties" yaml:"gameProperties"`
	// A set of custom game session properties, formatted as a single string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#game_session_data GameliftMatchmakingConfiguration#game_session_data}
	GameSessionData *string `field:"optional" json:"gameSessionData" yaml:"gameSessionData"`
	// The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#game_session_queue_arns GameliftMatchmakingConfiguration#game_session_queue_arns}
	GameSessionQueueArns *[]*string `field:"optional" json:"gameSessionQueueArns" yaml:"gameSessionQueueArns"`
	// An SNS topic ARN that is set up to receive matchmaking notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#notification_target GameliftMatchmakingConfiguration#notification_target}
	NotificationTarget *string `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#rule_set_arn GameliftMatchmakingConfiguration#rule_set_arn}
	RuleSetArn *string `field:"optional" json:"ruleSetArn" yaml:"ruleSetArn"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_matchmaking_configuration#tags GameliftMatchmakingConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


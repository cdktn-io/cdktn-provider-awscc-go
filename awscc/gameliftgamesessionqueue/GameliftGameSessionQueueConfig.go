// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgamesessionqueue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftGameSessionQueueConfig struct {
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
	// A descriptive label that is associated with game session queue. Queue names must be unique within each Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#name GameliftGameSessionQueue#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Information that is added to all events that are related to this game session queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#custom_event_data GameliftGameSessionQueue#custom_event_data}
	CustomEventData *string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#destinations GameliftGameSessionQueue#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// A list of locations where a queue is allowed to place new game sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#filter_configuration GameliftGameSessionQueue#filter_configuration}
	FilterConfiguration *GameliftGameSessionQueueFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// An SNS topic ARN that is set up to receive game session placement notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#notification_target GameliftGameSessionQueue#notification_target}
	NotificationTarget *string `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// A set of policies that act as a sliding cap on player latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#player_latency_policies GameliftGameSessionQueue#player_latency_policies}
	PlayerLatencyPolicies interface{} `field:"optional" json:"playerLatencyPolicies" yaml:"playerLatencyPolicies"`
	// Custom settings to use when prioritizing destinations and locations for game session placements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#priority_configuration GameliftGameSessionQueue#priority_configuration}
	PriorityConfiguration *GameliftGameSessionQueuePriorityConfiguration `field:"optional" json:"priorityConfiguration" yaml:"priorityConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#tags GameliftGameSessionQueue#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The maximum time, in seconds, that a new game session placement request remains in the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_game_session_queue#timeout_in_seconds GameliftGameSessionQueue#timeout_in_seconds}
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}


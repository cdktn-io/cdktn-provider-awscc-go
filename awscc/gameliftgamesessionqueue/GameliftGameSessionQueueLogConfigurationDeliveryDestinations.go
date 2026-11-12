// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgamesessionqueue


type GameliftGameSessionQueueLogConfigurationDeliveryDestinations struct {
	// The Amazon Resource Name (ARN) of the resource that receives the placement logs.
	//
	// The resource type must match LogDestination (a CloudWatch Logs log group, S3 bucket, or Firehose delivery stream). Required for FIREHOSE; optional for CLOUDWATCH and S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/gamelift_game_session_queue#destination_arn GameliftGameSessionQueue#destination_arn}
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// The type of service that receives the game session queue's placement logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/gamelift_game_session_queue#log_destination GameliftGameSessionQueue#log_destination}
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectroutingprofile


type ConnectRoutingProfileManualAssignmentQueueConfigsQueueReference struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_routing_profile#channel ConnectRoutingProfile#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The Amazon Resource Name (ARN) for the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_routing_profile#queue_arn ConnectRoutingProfile#queue_arn}
	QueueArn *string `field:"optional" json:"queueArn" yaml:"queueArn"`
}


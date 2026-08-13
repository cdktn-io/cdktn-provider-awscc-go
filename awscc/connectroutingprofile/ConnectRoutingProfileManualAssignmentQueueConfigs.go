// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectroutingprofile


type ConnectRoutingProfileManualAssignmentQueueConfigs struct {
	// Contains the channel and queue identifier for a routing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_routing_profile#queue_reference ConnectRoutingProfile#queue_reference}
	QueueReference *ConnectRoutingProfileManualAssignmentQueueConfigsQueueReference `field:"optional" json:"queueReference" yaml:"queueReference"`
}


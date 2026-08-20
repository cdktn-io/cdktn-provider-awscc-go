// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectroutingprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectRoutingProfileConfig struct {
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
	// The identifier of the default outbound queue for this routing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#default_outbound_queue_arn ConnectRoutingProfile#default_outbound_queue_arn}
	DefaultOutboundQueueArn *string `field:"required" json:"defaultOutboundQueueArn" yaml:"defaultOutboundQueueArn"`
	// The description of the routing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#description ConnectRoutingProfile#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#instance_arn ConnectRoutingProfile#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#media_concurrencies ConnectRoutingProfile#media_concurrencies}
	MediaConcurrencies interface{} `field:"required" json:"mediaConcurrencies" yaml:"mediaConcurrencies"`
	// The name of the routing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#name ConnectRoutingProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether agents with this routing profile will have their routing order calculated based on longest idle time or time since their last inbound contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#agent_availability_timer ConnectRoutingProfile#agent_availability_timer}
	AgentAvailabilityTimer *string `field:"optional" json:"agentAvailabilityTimer" yaml:"agentAvailabilityTimer"`
	// The manual assignment queues to associate with this routing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#manual_assignment_queue_configs ConnectRoutingProfile#manual_assignment_queue_configs}
	ManualAssignmentQueueConfigs interface{} `field:"optional" json:"manualAssignmentQueueConfigs" yaml:"manualAssignmentQueueConfigs"`
	// The queues to associate with this routing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#queue_configs ConnectRoutingProfile#queue_configs}
	QueueConfigs interface{} `field:"optional" json:"queueConfigs" yaml:"queueConfigs"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_routing_profile#tags ConnectRoutingProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


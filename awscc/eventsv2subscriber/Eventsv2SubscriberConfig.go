// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2SubscriberConfig struct {
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
	// The ARN of the event bus this subscriber belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#event_bus_arn Eventsv2Subscriber#event_bus_arn}
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
	// Configuration for how the subscriber invokes its target, including the target ARN, the IAM role used to invoke it, and, optionally, the target-specific parameters object that matches the target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#invoke_configuration Eventsv2Subscriber#invoke_configuration}
	InvokeConfiguration *Eventsv2SubscriberInvokeConfiguration `field:"required" json:"invokeConfiguration" yaml:"invokeConfiguration"`
	// The name of the subscriber.
	//
	// The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#name Eventsv2Subscriber#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration for batching events into a single delivery to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#batch_configuration Eventsv2Subscriber#batch_configuration}
	BatchConfiguration *Eventsv2SubscriberBatchConfiguration `field:"optional" json:"batchConfiguration" yaml:"batchConfiguration"`
	// A description of the subscriber. Control characters and Unicode line separators are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#description Eventsv2Subscriber#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration for filtering which events are delivered to the target. An event must match every filter to be delivered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#filter_configuration Eventsv2Subscriber#filter_configuration}
	FilterConfiguration *Eventsv2SubscriberFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// Delivery logging configuration for the subscriber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#log_configuration Eventsv2Subscriber#log_configuration}
	LogConfiguration *Eventsv2SubscriberLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The destination for events that could not be delivered to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#on_failure_configuration Eventsv2Subscriber#on_failure_configuration}
	OnFailureConfiguration *Eventsv2SubscriberOnFailureConfiguration `field:"optional" json:"onFailureConfiguration" yaml:"onFailureConfiguration"`
	// The point in time to start delivering events from. Used when StartingPosition is POINT_IN_TIME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#point_in_time_configuration Eventsv2Subscriber#point_in_time_configuration}
	PointInTimeConfiguration *Eventsv2SubscriberPointInTimeConfiguration `field:"optional" json:"pointInTimeConfiguration" yaml:"pointInTimeConfiguration"`
	// Resume-time control, never returned by the service.
	//
	// Applied only when an update transitions State from STOPPED to RUNNING: LAST_PROCESSED (default) resumes from the last processed event, LATEST skips to the newest. Ignored on create and on any update that does not perform that transition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#resume_position Eventsv2Subscriber#resume_position}
	ResumePosition *string `field:"optional" json:"resumePosition" yaml:"resumePosition"`
	// The retry policy for failed deliveries to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#retry_policy Eventsv2Subscriber#retry_policy}
	RetryPolicy *Eventsv2SubscriberRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Where the subscriber starts reading events: LATEST starts from the newest events;
	//
	// POINT_IN_TIME starts from the point specified in PointInTimeConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#starting_position Eventsv2Subscriber#starting_position}
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// The run state of the subscriber.
	//
	// Events are delivered only while the state is RUNNING. Setting the state to STOPPED pauses delivery. When an update sets a stopped subscriber back to RUNNING, ResumePosition controls where delivery resumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#state Eventsv2Subscriber#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags assigned to the subscriber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#tags Eventsv2Subscriber#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Configuration for transforming events before delivery to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#transformer Eventsv2Subscriber#transformer}
	Transformer *Eventsv2SubscriberTransformer `field:"optional" json:"transformer" yaml:"transformer"`
	// The delivery ordering mode of the subscriber.
	//
	// FIFO delivers events in order within an event group; UNORDERED delivers without an ordering guarantee.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#type Eventsv2Subscriber#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


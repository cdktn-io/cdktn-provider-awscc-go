// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventseventbus

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EventsEventBusConfig struct {
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
	// The name of the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_event_bus#name EventsEventBus#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dead Letter Queue for the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_event_bus#dead_letter_config EventsEventBus#dead_letter_config}
	DeadLetterConfig *EventsEventBusDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The description of the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_event_bus#description EventsEventBus#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_event_bus#event_source_name EventsEventBus#event_source_name}
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
	// Kms Key Identifier used to encrypt events at rest in the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_event_bus#kms_key_identifier EventsEventBus#kms_key_identifier}
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The logging configuration settings for vended logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_event_bus#log_config EventsEventBus#log_config}
	LogConfig *EventsEventBusLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// A JSON string that describes the permission policy statement for the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_event_bus#policy EventsEventBus#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Any tags assigned to the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_event_bus#tags EventsEventBus#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


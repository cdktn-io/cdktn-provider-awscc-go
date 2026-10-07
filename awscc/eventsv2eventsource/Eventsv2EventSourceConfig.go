// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventsource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2EventSourceConfig struct {
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
	// The event source configuration. Specify exactly one of AwsServiceEventsConfiguration or PartnerEventsConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#configuration Eventsv2EventSource#configuration}
	Configuration *Eventsv2EventSourceConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The ARN of the custom event bus the event source forwards onto.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#event_bus_arn Eventsv2EventSource#event_bus_arn}
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
	// The name of the event source.
	//
	// The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'. Names cannot begin with the reserved aws. prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#name Eventsv2EventSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the event source. Control characters and Unicode line separators are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#description Eventsv2EventSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags assigned to the event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_event_source#tags Eventsv2EventSource#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofileseventtrigger

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesEventTriggerConfig struct {
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
	// The unique name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_event_trigger#domain_name CustomerprofilesEventTrigger#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// A list of conditions that determine when an event should trigger the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_event_trigger#event_trigger_conditions CustomerprofilesEventTrigger#event_trigger_conditions}
	EventTriggerConditions interface{} `field:"required" json:"eventTriggerConditions" yaml:"eventTriggerConditions"`
	// The unique name of the event trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_event_trigger#event_trigger_name CustomerprofilesEventTrigger#event_trigger_name}
	EventTriggerName *string `field:"required" json:"eventTriggerName" yaml:"eventTriggerName"`
	// The unique name of the object type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_event_trigger#object_type_name CustomerprofilesEventTrigger#object_type_name}
	ObjectTypeName *string `field:"required" json:"objectTypeName" yaml:"objectTypeName"`
	// The description of the event trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_event_trigger#description CustomerprofilesEventTrigger#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_event_trigger#event_trigger_limits CustomerprofilesEventTrigger#event_trigger_limits}
	EventTriggerLimits *CustomerprofilesEventTriggerEventTriggerLimits `field:"optional" json:"eventTriggerLimits" yaml:"eventTriggerLimits"`
	// The destination is triggered only for profiles that meet the criteria of a segment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_event_trigger#segment_filter CustomerprofilesEventTrigger#segment_filter}
	SegmentFilter *string `field:"optional" json:"segmentFilter" yaml:"segmentFilter"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_event_trigger#tags CustomerprofilesEventTrigger#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


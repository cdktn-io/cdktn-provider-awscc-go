// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveeventbridgeruletemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveEventBridgeRuleTemplateConfig struct {
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
	// The type of event to match with the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_event_bridge_rule_template#event_type MedialiveEventBridgeRuleTemplate#event_type}
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_event_bridge_rule_template#name MedialiveEventBridgeRuleTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A resource's optional description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_event_bridge_rule_template#description MedialiveEventBridgeRuleTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Placeholder documentation for __listOfEventBridgeRuleTemplateTarget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_event_bridge_rule_template#event_targets MedialiveEventBridgeRuleTemplate#event_targets}
	EventTargets interface{} `field:"optional" json:"eventTargets" yaml:"eventTargets"`
	// An eventbridge rule template group's identifier. Can be either be its id or current name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_event_bridge_rule_template#group_identifier MedialiveEventBridgeRuleTemplate#group_identifier}
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
	// Represents the tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_event_bridge_rule_template#tags MedialiveEventBridgeRuleTemplate#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


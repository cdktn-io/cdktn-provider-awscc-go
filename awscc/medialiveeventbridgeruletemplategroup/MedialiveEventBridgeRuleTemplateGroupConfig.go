// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveeventbridgeruletemplategroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveEventBridgeRuleTemplateGroupConfig struct {
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
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/medialive_event_bridge_rule_template_group#name MedialiveEventBridgeRuleTemplateGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A resource's optional description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/medialive_event_bridge_rule_template_group#description MedialiveEventBridgeRuleTemplateGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Represents the tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/medialive_event_bridge_rule_template_group#tags MedialiveEventBridgeRuleTemplateGroup#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


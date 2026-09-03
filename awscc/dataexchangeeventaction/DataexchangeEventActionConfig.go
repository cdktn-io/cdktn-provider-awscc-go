// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataexchangeEventActionConfig struct {
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
	// What occurs after a certain event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dataexchange_event_action#action DataexchangeEventAction#action}
	Action *DataexchangeEventActionAction `field:"required" json:"action" yaml:"action"`
	// What occurs to start an action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dataexchange_event_action#event DataexchangeEventAction#event}
	Event *DataexchangeEventActionEvent `field:"required" json:"event" yaml:"event"`
	// The tags for the event action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dataexchange_event_action#tags DataexchangeEventAction#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


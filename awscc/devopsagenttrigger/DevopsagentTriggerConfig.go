// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagenttrigger

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentTriggerConfig struct {
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
	// The action to perform when the trigger fires. A JSON object containing actionType and task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_trigger#action DevopsagentTrigger#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The unique identifier of the parent Agent Space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_trigger#agent_space_id DevopsagentTrigger#agent_space_id}
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The condition that causes the trigger to fire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_trigger#condition DevopsagentTrigger#condition}
	Condition *DevopsagentTriggerCondition `field:"required" json:"condition" yaml:"condition"`
	// The type of trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_trigger#type DevopsagentTrigger#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The status of the trigger. Active triggers fire on schedule; Inactive triggers are paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_trigger#status DevopsagentTrigger#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


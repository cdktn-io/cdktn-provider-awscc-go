// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetrigger

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTriggerConfig struct {
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
	// The actions initiated by this trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#actions GlueTrigger#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The type of trigger that this is.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#type GlueTrigger#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of this trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#description GlueTrigger#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#event_batching_condition GlueTrigger#event_batching_condition}
	EventBatchingCondition *GlueTriggerEventBatchingCondition `field:"optional" json:"eventBatchingCondition" yaml:"eventBatchingCondition"`
	// The name of the trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#name GlueTrigger#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The predicate of this trigger, which defines when it will fire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#predicate GlueTrigger#predicate}
	Predicate *GlueTriggerPredicate `field:"optional" json:"predicate" yaml:"predicate"`
	// A cron expression used to specify the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#schedule GlueTrigger#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#start_on_creation GlueTrigger#start_on_creation}
	StartOnCreation interface{} `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
	// The tags to use with this trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#tags GlueTrigger#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workflow associated with the trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_trigger#workflow_name GlueTrigger#workflow_name}
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}


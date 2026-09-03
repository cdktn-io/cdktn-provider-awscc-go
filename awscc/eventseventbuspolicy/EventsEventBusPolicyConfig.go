// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventseventbuspolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EventsEventBusPolicyConfig struct {
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
	// An identifier string for the external account that you are granting permissions to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#statement_id EventsEventBusPolicy#statement_id}
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
	// The action that you are enabling the other account to perform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#action EventsEventBusPolicy#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// This parameter enables you to limit the permission to accounts that fulfill a certain condition, such as being a member of a certain AWS organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#condition EventsEventBusPolicy#condition}
	Condition *EventsEventBusPolicyCondition `field:"optional" json:"condition" yaml:"condition"`
	// The name of the event bus associated with the rule.
	//
	// If you omit this, the default event bus is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#event_bus_name EventsEventBusPolicy#event_bus_name}
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The 12-digit AWS account ID that you are permitting to put events to your default event bus.
	//
	// Specify "*" to permit any account to put events to your default event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#principal EventsEventBusPolicy#principal}
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// A JSON string that describes the permission policy statement.
	//
	// You can include a Policy parameter in the request instead of using the StatementId, Action, Principal, or Condition parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/events_event_bus_policy#statement EventsEventBusPolicy#statement}
	Statement *string `field:"optional" json:"statement" yaml:"statement"`
}


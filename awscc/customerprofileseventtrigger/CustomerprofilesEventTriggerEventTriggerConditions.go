// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofileseventtrigger


type CustomerprofilesEventTriggerEventTriggerConditions struct {
	// A list of dimensions to be evaluated for the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_event_trigger#event_trigger_dimensions CustomerprofilesEventTrigger#event_trigger_dimensions}
	EventTriggerDimensions interface{} `field:"required" json:"eventTriggerDimensions" yaml:"eventTriggerDimensions"`
	// The operator used to combine multiple dimensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_event_trigger#logical_operator CustomerprofilesEventTrigger#logical_operator}
	LogicalOperator *string `field:"required" json:"logicalOperator" yaml:"logicalOperator"`
}


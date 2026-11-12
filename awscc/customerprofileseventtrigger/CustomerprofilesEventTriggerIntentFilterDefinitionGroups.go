// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofileseventtrigger


type CustomerprofilesEventTriggerIntentFilterDefinitionGroups struct {
	// A list of intent filter conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_event_trigger#conditions CustomerprofilesEventTrigger#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The operator used to combine multiple dimensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_event_trigger#logical_operator CustomerprofilesEventTrigger#logical_operator}
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
}


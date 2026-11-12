// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofileseventtrigger


type CustomerprofilesEventTriggerIntentFilterDefinitionGroupsConditions struct {
	// The operator used to compare the predicted intent against a list of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_event_trigger#comparison_operator CustomerprofilesEventTrigger#comparison_operator}
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Identifies a predicted intent by domain and type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_event_trigger#intent CustomerprofilesEventTrigger#intent}
	Intent *CustomerprofilesEventTriggerIntentFilterDefinitionGroupsConditionsIntent `field:"optional" json:"intent" yaml:"intent"`
	// A list of intent values used for comparison.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_event_trigger#values CustomerprofilesEventTrigger#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


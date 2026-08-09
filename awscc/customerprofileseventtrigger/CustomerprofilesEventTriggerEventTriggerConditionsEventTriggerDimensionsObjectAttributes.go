// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofileseventtrigger


type CustomerprofilesEventTriggerEventTriggerConditionsEventTriggerDimensionsObjectAttributes struct {
	// The operator used to compare an attribute against a list of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_event_trigger#comparison_operator CustomerprofilesEventTrigger#comparison_operator}
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// A list of attribute values used for comparison.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_event_trigger#values CustomerprofilesEventTrigger#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// A field defined within an object type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_event_trigger#field_name CustomerprofilesEventTrigger#field_name}
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// An attribute contained within a source object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_event_trigger#source CustomerprofilesEventTrigger#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}


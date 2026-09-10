// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetrigger


type GlueTriggerPredicate struct {
	// A list of the conditions that determine when the trigger will fire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#conditions GlueTrigger#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// An optional field if only one condition is listed. If multiple conditions are listed, then this field is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#logical GlueTrigger#logical}
	Logical *string `field:"optional" json:"logical" yaml:"logical"`
}


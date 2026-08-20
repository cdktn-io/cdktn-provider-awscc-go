// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsplan


type SsmcontactsPlanStagesTargetsContactTargetInfo struct {
	// The Amazon Resource Name (ARN) of the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmcontacts_plan#contact_id SsmcontactsPlan#contact_id}
	ContactId *string `field:"optional" json:"contactId" yaml:"contactId"`
	// A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmcontacts_plan#is_essential SsmcontactsPlan#is_essential}
	IsEssential interface{} `field:"optional" json:"isEssential" yaml:"isEssential"`
}


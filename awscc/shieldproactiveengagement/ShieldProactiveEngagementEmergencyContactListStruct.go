// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shieldproactiveengagement


type ShieldProactiveEngagementEmergencyContactListStruct struct {
	// The email address for the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/shield_proactive_engagement#email_address ShieldProactiveEngagement#email_address}
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// Additional notes regarding the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/shield_proactive_engagement#contact_notes ShieldProactiveEngagement#contact_notes}
	ContactNotes *string `field:"optional" json:"contactNotes" yaml:"contactNotes"`
	// The phone number for the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/shield_proactive_engagement#phone_number ShieldProactiveEngagement#phone_number}
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}


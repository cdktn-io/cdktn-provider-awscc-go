// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactscontact


type SsmcontactsContactPlan struct {
	// The time to wait until beginning the next stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_contact#duration_in_minutes SsmcontactsContact#duration_in_minutes}
	DurationInMinutes *float64 `field:"optional" json:"durationInMinutes" yaml:"durationInMinutes"`
	// List of Rotation Ids to associate with Contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_contact#rotation_ids SsmcontactsContact#rotation_ids}
	RotationIds *[]*string `field:"optional" json:"rotationIds" yaml:"rotationIds"`
	// The contacts or contact methods that the escalation plan or engagement plan is engaging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_contact#targets SsmcontactsContact#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}


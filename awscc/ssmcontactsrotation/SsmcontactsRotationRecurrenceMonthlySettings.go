// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation


type SsmcontactsRotationRecurrenceMonthlySettings struct {
	// The day of the month when monthly recurring on-call rotations begin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_rotation#day_of_month SsmcontactsRotation#day_of_month}
	DayOfMonth *float64 `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_rotation#hand_off_time SsmcontactsRotation#hand_off_time}
	HandOffTime *string `field:"optional" json:"handOffTime" yaml:"handOffTime"`
}


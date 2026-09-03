// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation


type SsmcontactsRotationRecurrenceShiftCoverages struct {
	// Information about when an on-call shift begins and ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_rotation#coverage_times SsmcontactsRotation#coverage_times}
	CoverageTimes interface{} `field:"optional" json:"coverageTimes" yaml:"coverageTimes"`
	// The day of the week when weekly recurring on-call shift rotations begin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_rotation#day_of_week SsmcontactsRotation#day_of_week}
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
}


// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation


type SsmcontactsRotationRecurrenceShiftCoveragesCoverageTimes struct {
	// Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmcontacts_rotation#end_time SsmcontactsRotation#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmcontacts_rotation#start_time SsmcontactsRotation#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}


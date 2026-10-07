// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramScheduleConfigurationTransition struct {
	// The duration of the live program in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#duration_millis MediatailorProgram#duration_millis}
	DurationMillis *float64 `field:"optional" json:"durationMillis" yaml:"durationMillis"`
	// The position where this program will be inserted relative to the RelativePosition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#relative_position MediatailorProgram#relative_position}
	RelativePosition *string `field:"optional" json:"relativePosition" yaml:"relativePosition"`
	// The name of the program that this program will be inserted next to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#relative_program MediatailorProgram#relative_program}
	RelativeProgram *string `field:"optional" json:"relativeProgram" yaml:"relativeProgram"`
	// The date and time that the program is scheduled to start, in epoch milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#scheduled_start_time_millis MediatailorProgram#scheduled_start_time_millis}
	ScheduledStartTimeMillis *float64 `field:"optional" json:"scheduledStartTimeMillis" yaml:"scheduledStartTimeMillis"`
	// Defines when the program plays in the schedule. You can set the value to ABSOLUTE or RELATIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#type MediatailorProgram#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


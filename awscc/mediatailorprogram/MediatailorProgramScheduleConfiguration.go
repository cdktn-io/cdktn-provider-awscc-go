// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramScheduleConfiguration struct {
	// Clip range configuration for the VOD source associated with the program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#clip_range MediatailorProgram#clip_range}
	ClipRange *MediatailorProgramScheduleConfigurationClipRange `field:"optional" json:"clipRange" yaml:"clipRange"`
	// Program transition configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#transition MediatailorProgram#transition}
	Transition *MediatailorProgramScheduleConfigurationTransition `field:"optional" json:"transition" yaml:"transition"`
}


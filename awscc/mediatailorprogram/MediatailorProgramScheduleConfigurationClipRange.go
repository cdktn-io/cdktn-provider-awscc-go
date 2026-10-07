// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramScheduleConfigurationClipRange struct {
	// The end offset of the clip range, in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#end_offset_millis MediatailorProgram#end_offset_millis}
	EndOffsetMillis *float64 `field:"optional" json:"endOffsetMillis" yaml:"endOffsetMillis"`
	// The start offset of the clip range, in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#start_offset_millis MediatailorProgram#start_offset_millis}
	StartOffsetMillis *float64 `field:"optional" json:"startOffsetMillis" yaml:"startOffsetMillis"`
}


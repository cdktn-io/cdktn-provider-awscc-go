// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramAudienceMediaAlternateMedia struct {
	// Ad break configuration parameters defined in AlternateMedia.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#ad_breaks MediatailorProgram#ad_breaks}
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// Clip range configuration for the VOD source associated with the program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#clip_range MediatailorProgram#clip_range}
	ClipRange *MediatailorProgramAudienceMediaAlternateMediaClipRange `field:"optional" json:"clipRange" yaml:"clipRange"`
	// The duration of the alternateMedia in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#duration_millis MediatailorProgram#duration_millis}
	DurationMillis *float64 `field:"optional" json:"durationMillis" yaml:"durationMillis"`
	// The name of the live source for alternateMedia.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#live_source_name MediatailorProgram#live_source_name}
	LiveSourceName *string `field:"optional" json:"liveSourceName" yaml:"liveSourceName"`
	// The date and time that the alternateMedia is scheduled to start, in epoch milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#scheduled_start_time_millis MediatailorProgram#scheduled_start_time_millis}
	ScheduledStartTimeMillis *float64 `field:"optional" json:"scheduledStartTimeMillis" yaml:"scheduledStartTimeMillis"`
	// The name of the source location for alternateMedia.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#source_location_name MediatailorProgram#source_location_name}
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The name of the VOD source for alternateMedia.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#vod_source_name MediatailorProgram#vod_source_name}
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}


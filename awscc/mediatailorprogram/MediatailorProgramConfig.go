// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorProgramConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the channel for this Program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#channel_name MediatailorProgram#channel_name}
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The name of the Program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#program_name MediatailorProgram#program_name}
	ProgramName *string `field:"required" json:"programName" yaml:"programName"`
	// The name of the source location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#source_location_name MediatailorProgram#source_location_name}
	SourceLocationName *string `field:"required" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The ad break configuration settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#ad_breaks MediatailorProgram#ad_breaks}
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// The list of AudienceMedia defined in program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#audience_media MediatailorProgram#audience_media}
	AudienceMedia interface{} `field:"optional" json:"audienceMedia" yaml:"audienceMedia"`
	// The name of the LiveSource for this Program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#live_source_name MediatailorProgram#live_source_name}
	LiveSourceName *string `field:"optional" json:"liveSourceName" yaml:"liveSourceName"`
	// The schedule configuration settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#schedule_configuration MediatailorProgram#schedule_configuration}
	ScheduleConfiguration *MediatailorProgramScheduleConfiguration `field:"optional" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The name that's used to refer to a VOD source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#vod_source_name MediatailorProgram#vod_source_name}
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}


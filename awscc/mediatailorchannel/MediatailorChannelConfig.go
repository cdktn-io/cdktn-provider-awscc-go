// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorchannel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorChannelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#channel_name MediatailorChannel#channel_name}.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// <p>The channel's output properties.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#outputs MediatailorChannel#outputs}
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#playback_mode MediatailorChannel#playback_mode}.
	PlaybackMode *string `field:"required" json:"playbackMode" yaml:"playbackMode"`
	// <p>The list of audiences defined in channel.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#audiences MediatailorChannel#audiences}
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// <p>Slate VOD source configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#filler_slate MediatailorChannel#filler_slate}
	FillerSlate *MediatailorChannelFillerSlate `field:"optional" json:"fillerSlate" yaml:"fillerSlate"`
	// <p>The log configuration for the channel.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#log_configuration MediatailorChannel#log_configuration}
	LogConfiguration *MediatailorChannelLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The tags to assign to the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#tags MediatailorChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#tier MediatailorChannel#tier}.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// <p>The configuration for time-shifted viewing.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#time_shift_configuration MediatailorChannel#time_shift_configuration}
	TimeShiftConfiguration *MediatailorChannelTimeShiftConfiguration `field:"optional" json:"timeShiftConfiguration" yaml:"timeShiftConfiguration"`
}


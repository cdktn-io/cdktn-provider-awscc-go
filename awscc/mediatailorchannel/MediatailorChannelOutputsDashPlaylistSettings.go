// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorchannel


type MediatailorChannelOutputsDashPlaylistSettings struct {
	// <p>The total duration (in seconds) of each manifest. Minimum value: <code>30</code> seconds. Maximum value: <code>3600</code> seconds.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_channel#manifest_window_seconds MediatailorChannel#manifest_window_seconds}
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// <p>Minimum amount of content (measured in seconds) that a player must keep available in the buffer.
	//
	// Minimum value: <code>2</code> seconds. Maximum value: <code>60</code> seconds.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_channel#min_buffer_time_seconds MediatailorChannel#min_buffer_time_seconds}
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// <p>Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.
	//
	// Minimum value: <code>2</code> seconds. Maximum value: <code>60</code> seconds.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_channel#min_update_period_seconds MediatailorChannel#min_update_period_seconds}
	MinUpdatePeriodSeconds *float64 `field:"optional" json:"minUpdatePeriodSeconds" yaml:"minUpdatePeriodSeconds"`
	// <p>Amount of time (in seconds) that the player should be from the live point at the end of the manifest.
	//
	// Minimum value: <code>2</code> seconds. Maximum value: <code>60</code> seconds.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediatailor_channel#suggested_presentation_delay_seconds MediatailorChannel#suggested_presentation_delay_seconds}
	SuggestedPresentationDelaySeconds *float64 `field:"optional" json:"suggestedPresentationDelaySeconds" yaml:"suggestedPresentationDelaySeconds"`
}


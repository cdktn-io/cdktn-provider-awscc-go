// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorchannel


type MediatailorChannelTimeShiftConfiguration struct {
	// <p>The maximum time delay for time-shifted viewing.
	//
	// The minimum allowed maximum time delay is 0 seconds, and the maximum allowed maximum time delay is 21600 seconds (6 hours).</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_channel#max_time_delay_seconds MediatailorChannel#max_time_delay_seconds}
	MaxTimeDelaySeconds *float64 `field:"optional" json:"maxTimeDelaySeconds" yaml:"maxTimeDelaySeconds"`
}

